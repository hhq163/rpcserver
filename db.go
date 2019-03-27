package main

import (
	"database/sql"
	"mangos/core/slog"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/mgo.v2"
)

var mgoSession *mgo.Session

//workers协程使用的连接池
var mysqlDB *sql.DB

//InitDB init db
func InitDB() {
	var err error
	mgoSession, err = mgo.Dial(Cfg.MongoURL)
	if err != nil {
		slog.Fatal(err)
	}
	mgoSession.SetPoolLimit(Cfg.MongoThreads)
	mysqlDB, err = sql.Open("mysql", Cfg.MySqlURL)
	if err != nil {
		slog.Fatal(err)
	}
	mysqlDB.SetMaxOpenConns(Cfg.MySqlThreads)
	mysqlDB.SetMaxIdleConns(Cfg.MySqlThreads)

}
