package main

import (
	"mangos/core/slog"
	"rpcserver/util"
)

var mysqlWorkList, mongoWorkList *util.WorkList

//StartWorkers start workers
func StartWorkers() {
	if Cfg.MongoThreads <= 0 {
		slog.Fatal("Cfg.MongoThreads <= 0")
	}

	mysqlWorkList = util.NewWorkList(Cfg.MySqlThreads)
	mongoWorkList = util.NewWorkList(Cfg.MongoThreads)

}

//StopWorkers stop workers
func StopWorkers() {
	mysqlWorkList.Close()
	mongoWorkList.Close()
}

func pushmysql(f func()) {
	mysqlWorkList.Push(f)
}

func pushmongo(f func()) {
	mongoWorkList.Push(f)
}
