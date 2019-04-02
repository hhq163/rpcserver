package main

import (
	"math/rand"
	"rpcserver/protocol"
	"rpcserver/slog"
	"time"
)

var chanLen = 10000
var workNumber = 30

//邮件消息
type sqlMsg struct {
	AddTime  time.Time
	HallID   int32
	AgentID  int32
	UserID   int64
	HallName string
	UserName string
	Amount   float64
	Itype    int32
	TimeOut  int //单位为秒
	Res      chan *protocol.CashOperResponse
}

var sqlTaskChan = make(chan *sqlMsg, chanLen)
var sqlTaskMap = make(map[int]chan *sqlMsg, workNumber)

func init() {
	go SqlMainTask()

	// work pool goroutine
	for i := 0; i < workNumber; i++ {
		taskChan := make(chan *sqlMsg, chanLen)
		sqlTaskMap[i] = taskChan // insert to task map
		go SqlWorkTask(taskChan)
	}
}

//sql消息发送主协程
func SqlMainTask() {
	for msg := range sqlTaskChan {
		rand.Seed(time.Now().Unix())
		index := rand.Intn(workNumber)
		//slog.Info("EmailMainTask, index=", index)

		if ok := sqlTaskMap[index]; ok != nil {
			sqlTaskMap[index] <- msg
		}
	}
}

//sql工作协程
func SqlWorkTask(taskChan chan *sqlMsg) {

	for msg := range taskChan {
		now := time.Now().Unix()
		//消息超时
		if msg.AddTime.Unix()+int64(msg.TimeOut) < now {
			msg.Res <- &protocol.CashOperResponse{
				ResultCode: 1,
				Desc:       "out of time",
				Restult: &protocol.Result{
					Amount:  msg.Amount,
					OrderSn: "",
				},
			}
			return
		}

		result, err := mysqlDB.Exec("UPDATE lb_user_1 SET money=money + ? WHERE hall_id=? AND uid=? LIMIT 1", msg.Amount, msg.HallId, msg.UserId)
		if err != nil {
			slog.ErrorDB(err)
			msg.Res <- &protocol.CashOperResponse{
				ResultCode: 1,
				Desc:       "update record failed ",
				Restult: &protocol.Result{
					Amount:  msg.Amount,
					OrderSn: "",
				},
			}
			return
		}
		rows, err := result.RowsAffected()
		if err != nil {
			slog.ErrorDB(err)
		}
		if rows < 0 { //未修改成功
			slog.ErrorDB("update money failed ,hall_id=", msg.HallId, "user_id=", msg.UserId)
		}
	}
}
