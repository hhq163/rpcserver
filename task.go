package main

import (
	"math/rand"
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
	amount   float64
	Itype    int32
	TimeOut  chan int //单位为秒
}

var sqlTaskChan = make(chan *sqlMsg, chanLen)
var sqlTaskMap = make(map[int]chan *sqlMsg, workNumber)

func init() {
	go SqlMainTask()

	// work pool goroutine
	for i := 0; i < workNumber; i++ {
		taskChan := make(chan *sqlMsg, chanLen)
		sqlTaskMap[i] = taskChan // insert to task map
		go EmailWorkTask(taskChan)
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
func EmailWorkTask(taskChan chan *sqlMsg) {

	for task := range taskChan {

	}
}
