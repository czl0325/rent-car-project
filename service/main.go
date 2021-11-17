package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	trippb "service/proto/gen/go"
	trip "service/tripService"
)

func main() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("服务器启动失败，错误=%v", err)
	}
	s := grpc.NewServer()
	trippb.RegisterTripServiceServer(s, &trip.Service{})
	s.Serve(listener)
}
