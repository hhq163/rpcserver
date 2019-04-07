package main

import (
	"encoding/json"
	"fmt"
	"io"
	"rpcserver/protocol"
	"rpcserver/slog"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/mgo.v2/bson"
)

type User struct{}

var TableNameMap map[int32]string

//GetTableName 获取分表名称
func GetTableName(hall_id int32) string {
	var tableName string
	if tn, ok := TableNameMap[hall_id]; ok {
		tableName = tn
	} else {
		//从db中查询出来，存入内存中

	}
	tableName = "lb_user_1"
	return tableName
}

//现金操作接口
func (u *User) CashOpera(stream protocol.User_CashOperaServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if req.HallId <= 0 || req.UserId <= 0 {
			return fmt.Errorf("Param is wrong!")
		}

		slog.Info("UserId:", req.UserId, ",AgentId:", req.AgentId, ",HallId:", req.HallId, ",UserName:", req.UserName, ", Amount:", req.Amount, "Type:", req.Type)

		pushmongo(func() {
			mgo := mgoSession.Copy()
			defer mgo.Close()

			doc := bson.M{
				"order_sn":  "qxtestid",
				"uid":       req.UserId,
				"agent_id":  req.AgentId,
				"hall_id":   req.HallId,
				"user_name": req.UserName,
				"type":      req.Type,
				"amount":    req.Amount,
				"add_time":  time.Now(),
			}

			if err := mgo.DB("").C("cash_record_test").Insert(doc); err != nil {
				docjosn, _ := json.Marshal(doc)
				slog.ErrorDB("insert into cash_record failed err:", err, string(docjosn))

				res := &protocol.CashOperResponse{
					ResultCode: 1,
					Desc:       "insert into cash_record failed",
					Restult: &protocol.Result{
						Amount:  req.Amount,
						OrderSn: "",
					},
				}
				err = stream.Send(res)
				if err != nil {
					slog.Error("stream.Send error, hall_id=", req.HallId, "user_id=", req.UserId, ", err=", err)
					return
				}
				return
			}
			pushmysql(func() {
				// tablename := GetTableName(req.HallId)
				result, err := mysqlDB.Exec(fmt.Sprintf("UPDATE lb_user_1 SET money=money + ? WHERE hall_id=? AND uid=? LIMIT 1"), req.Amount, req.HallId, req.UserId)
				if err != nil {
					slog.ErrorDB(err)
					res := &protocol.CashOperResponse{
						ResultCode: 2,
						Desc:       "update money failed",
						Restult: &protocol.Result{
							Amount:  req.Amount,
							OrderSn: "",
						},
					}
					err = stream.Send(res)
					if err != nil {
						slog.Error("stream.Send error, hall_id=", req.HallId, "user_id=", req.UserId, ", err=", err)
						return
					}
					return
				}
				rows, err := result.RowsAffected()
				if err != nil {
					slog.ErrorDB(err)
				}
				if rows < 0 { //未修改成功
					slog.ErrorDB("update money failed ,hall_id=", req.HallId, "user_id=", req.UserId)
				}

				res := &protocol.CashOperResponse{
					ResultCode: 0,
					Desc:       "send success",
					Restult: &protocol.Result{
						Amount:  req.Amount,
						OrderSn: "",
					},
				}

				err = stream.Send(res)
				if err != nil {
					slog.Error("stream.Send error, hall_id=", req.HallId, "user_id=", req.UserId, ", err=", err)
					return
				}
				return
			})
			return
		})

	}
	return nil

}
