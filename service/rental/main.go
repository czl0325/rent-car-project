package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	rentalpb "service/rental/api/gen/v1"
	"service/rental/rental"
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
		myLogger.Error("连接mongodb失败,错误=%v\n", zap.Error(err))
	}
	myLogger.Info("输出=", zap.Any("mongo", mc))
	err = server.RunGRPCServer(&server.GRPCConfig{
		Name:         "rental",
		Addr:         ":9200",
		RegisterFunc: func(g *grpc.Server) {
			rentalpb.RegisterTripServiceServer(g, &rental.Service{

			})
		},
		Logger:       myLogger,
	})
	if err != nil {
		myLogger.Error("启动rental微服务失败,错误=%v", zap.Error(err))
	}
}
