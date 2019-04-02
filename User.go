package main

import (
	"context"
	"encoding/json"
	"fmt"
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
func (u *User) CashOpera(ctx context.Context, req *protocol.CashOperRequest) (*protocol.CashOperResponse, error) {
	if req.HallId <= 0 || req.UserId <= 0 {
		return &protocol.CashOperResponse{
			ResultCode: -1,
			Desc:       "hall_id or user_id <=0",
			Restult: &protocol.Result{
				Amount:  req.Amount,
				OrderSn: "",
			},
		}, fmt.Errorf("Param is wrong!")
	}
	// slog.Info("UserId:", req.UserId, ",AgentId:", req.AgentId, ",HallId:", req.HallId, ",UserName:", req.UserName, ", Amount:", req.Amount, "Type:", req.Type)

	pushmongo(func() {
		reqs := *req
		mgo := mgoSession.Copy()
		defer mgo.Close()

		doc := bson.M{
			"order_sn":  "qxtestid",
			"uid":       reqs.UserId,
			"agent_id":  reqs.AgentId,
			"hall_id":   reqs.HallId,
			"user_name": reqs.UserName,
			"type":      reqs.Type,
			"amount":    reqs.Amount,
			"add_time":  time.Now(),
		}

		if err := mgo.DB("").C("cash_record_test").Insert(doc); err != nil {
			docjosn, _ := json.Marshal(doc)
			slog.ErrorDB("cash_record failed err:", err, string(docjosn))
			return
		}
		pushmysql(func() {
			// tablename := GetTableName(reqs.HallId)
			result, err := mysqlDB.Exec(fmt.Sprintf("UPDATE lb_user_1 SET money=money + ? WHERE hall_id=? AND uid=? LIMIT 1"), reqs.Amount, reqs.HallId, reqs.UserId)
			if err != nil {
				slog.ErrorDB(err)
				return
			}
			rows, err := result.RowsAffected()
			if err != nil {
				slog.ErrorDB(err)
			}
			if rows < 0 { //未修改成功
				slog.ErrorDB("update money failed ,hall_id=", reqs.HallId, "user_id=", reqs.UserId)
			}

		})
	})

	return &protocol.CashOperResponse{
		ResultCode: 1,
		Desc:       "oper success",
		Restult: &protocol.Result{
			Amount:  req.Amount,
			OrderSn: "",
		},
	}, nil

}
