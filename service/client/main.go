package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	trippb "service/proto/gen/go"
)

func main() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("客户端启动失败,错误=%s\n", err)
	}
	client := trippb.NewTripServiceClient(conn)
	res, err := client.GetTrip(context.Background(), &trippb.GetTripRequest{Id: "1"})
	fmt.Println(res)
}
