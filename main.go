package main

import (
	"log"
	"mangos/logicsvr/config"
	"net"
	"net/http"
	"rpcserver/protocol"
	"rpcserver/slog"

	_ "net/http/pprof"

	"google.golang.org/grpc"
)

func main() {
	if config.Cfg.DebugMode {
		go profile()
	}
	slog.NewLog(Cfg.LogLevel, true, 10)
	slog.Info("rpcserver starting up")

	lis, err := net.Listen("tcp", Cfg.Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//初始化数据库连接池
	InitDB()
	StartWorkers()

	s := grpc.NewServer()
	protocol.RegisterUserServer(s, &User{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func profile() {
	log.Fatal(http.ListenAndServe("0.0.0.0:9912", nil))
}
