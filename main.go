package main

import (
	"log"
	"net"
	"rpcserver/protocol"
	"rpcserver/slog"

	"google.golang.org/grpc"
)

func main() {
	slog.NewLog(Cfg.LogLevel, true, 10)
	slog.Info("rpcserver starting up")

	lis, err := net.Listen("tcp", Cfg.Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Fatalf("start 0000")
	s := grpc.NewServer()
	protocol.RegisterUserServer(s, &User{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Fatalf("start 1111")
	//初始化数据库连接池
	InitDB()
	StartWorkers()
	log.Fatalf("start 2222")
}
