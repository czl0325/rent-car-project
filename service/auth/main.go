package main

import (
	"google.golang.org/grpc"
	"log"
	authpb "service/auth/api/gen/v1"
	"service/auth/auth"
	"service/shared/server"
)

func main() {
	myLogger, err := server.NewZapLogger()
	if err != nil {
		log.Fatalf("日志模块初始化失败,错误=%v", err)
	}
	server.RunGRPCServer(&server.GRPCConfig{
		Name:              "auth",
		Addr:              ":9100",
		AuthPublicKeyFile: "",
		RegisterFunc: func(g *grpc.Server) {
			authpb.RegisterAuthServiceServer(g, &auth.Service{})
		},
		Logger:            myLogger,
	})
}