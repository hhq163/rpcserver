package main

import (
	"context"
	"fmt"
	"rpcserver/protocol"
	"time"

	_ "github.com/go-sql-driver/mysql"
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

	ResChan := make(chan *protocol.CashOperResponse)
	cashMsg := &sqlMsg{
		AddTime:  time.Now(),
		HallID:   req.HallId,
		AgentID:  req.AgentId,
		UserID:   req.UserId,
		HallName: req.HallName,
		UserName: req.UserName,
		Amount:   req.Amount,
		Itype:    int32(req.Type),
		TimeOut:  5,
		ResChan:  ResChan,
	}
	sqlTaskChan <- cashMsg

	res := <-ResChan
	if res.ResultCode == 0 {
		return res, nil
	} else {
		return res, fmt.Errorf("err: %s", res.Desc)
	}

}
