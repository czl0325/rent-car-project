package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"log"
	"net/http"
	authpb "service/auth/api/gen/v1"
	rentalpb "service/rental/api/gen/v1"
	"service/shared/server"
)

func main() {
	myLogger, err := server.NewZapLogger()
	if err != nil {
		log.Fatalf("gateway模块日志启动失败,错误=%v", err)
	}

	c := context.Background()
	c, cancel := context.WithCancel(c)
	defer cancel()

	mux := runtime.NewServeMux(runtime.WithMarshalerOption(
		runtime.MIMEWildcard,
		&runtime.JSONPb{
			OrigName:     true,
			EnumsAsInts:  true,
			EmitDefaults: false,
		}))

	serverConfig := []struct {
		Name         string
		Addr         string
		RegisterFunc func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error
	}{
		{
			Name:         "auth",
			Addr:         ":9100",
			RegisterFunc: authpb.RegisterAuthServiceHandlerFromEndpoint,
		},
		{
			Name:         "rental",
			Addr:         ":9200",
			RegisterFunc: rentalpb.RegisterTripServiceHandlerFromEndpoint,
		},
	}

	for _, s := range serverConfig{
		err = s.RegisterFunc(c, mux, s.Addr, []grpc.DialOption{grpc.WithInsecure()})
		if err != nil {
			myLogger.Sugar().Fatalf("无法启动服务,名字=%s,错误=%v", s.Name, err)
		}
	}
	http.Handle("/", mux)
	myLogger.Sugar().Fatal(http.ListenAndServe(":9000", nil))
	myLogger.Sugar().Info("网关服务启动成功!")
}
