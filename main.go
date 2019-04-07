package main

import (
	"log"
	"mangos/logicsvr/config"
	"net"
	"net/http"
	"rpcserver/protocol"
	"rpcserver/slog"
	"time"

	"google.golang.org/grpc/keepalive"

	_ "net/http/pprof"

	"google.golang.org/grpc"
)

var kaep = keepalive.EnforcementPolicy{
	MinTime:             5 * time.Second, // If a client pings more than once every 5 seconds, terminate the connection
	PermitWithoutStream: true,            // Allow pings even when there are no active streams
}

var kasp = keepalive.ServerParameters{
	MaxConnectionIdle:     15 * time.Second, // If a client is idle for 15 seconds, send a GOAWAY
	MaxConnectionAge:      30 * time.Second, // If any connection is alive for more than 30 seconds, send a GOAWAY
	MaxConnectionAgeGrace: 5 * time.Second,  // Allow 5 seconds for pending RPCs to complete before forcibly closing connections
	Time:    5 * time.Second, // Ping the client if it is idle for 5 seconds to ensure the connection is still active
	Timeout: 1 * time.Second, // Wait 1 second for the ping ack before assuming the connection is dead
}

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

	s := grpc.NewServer(grpc.KeepaliveEnforcementPolicy(kaep), grpc.KeepaliveParams(kasp))
	protocol.RegisterUserServer(s, &User{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func profile() {
	log.Fatal(http.ListenAndServe("0.0.0.0:9912", nil))
}
