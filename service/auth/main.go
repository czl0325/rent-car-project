package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	authpb "service/auth/api/gen/v1"
	"service/auth/auth"
	"service/auth/dao"
	"service/shared/server"
)

func main() {
	myLogger, err := server.NewZapLogger()
	if err != nil {
		log.Fatalf("日志模块初始化失败,错误=%v", err)
	}

	c := context.Background()
	mc, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		myLogger.Error("日志模块初始化失败,错误=%v\n", zap.Error(err))
	}

	err = server.RunGRPCServer(&server.GRPCConfig{
		Name: "auth",
		Addr: ":9100",
		RegisterFunc: func(g *grpc.Server) {
			authpb.RegisterAuthServiceServer(g, &auth.Service{
				Mongo: dao.NewMongo(mc.Database("rentcar")),
			})
		},
		Logger: myLogger,
	})
	if err != nil {
		myLogger.Error("启动auth微服务失败,错误=%v", zap.Error(err))
	}
}
