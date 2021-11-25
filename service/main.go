package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	trippb "service/proto/gen/go"
	trip "service/tripService"
)

func main() {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.TimeKey = ""
	myLog, err := config.Build()
	if err != nil {
		log.Fatalf("日志启动失败，错误=%v", err)
	}

	go startGRPCGateway()
	listener, err := net.Listen("tcp", ":9081")
	if err != nil {
		myLog.Sugar().Fatalf("服务器启动失败，错误=%v", err)
	}
	s := grpc.NewServer()
	trippb.RegisterTripServiceServer(s, &trip.Service{})
	err = s.Serve(listener)
	if err != nil {
		myLog.Sugar().Fatalf("服务器启动失败，错误=%v", err)
	}
}


func startGRPCGateway() {
	c := context.Background()
	c, cancel := context.WithCancel(c)
	defer cancel()
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: false,
	}))
	err := trippb.RegisterTripServiceHandlerFromEndpoint(c, mux, ":9100", []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatalf("启动grpc gateway失败。错误=%v\n", err)
	}
	err = http.ListenAndServe(":9000", mux)
	if err != nil {
		log.Fatalf("启动http服务失败。错误=%v\n", err)
	}
}