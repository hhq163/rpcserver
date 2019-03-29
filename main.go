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
	slog.Info("server starting up version:", version)

	lis, err := net.Listen("tcp", Cfg.TcpPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	protocol.RegisterUserServer(s, &User{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	//初始化数据库连接池
	InitDB()
	StartWorkers()

}
