package server

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

type GRPCConfig struct {
	Name              string
	Addr              string
	AuthPublicKeyFile string
	RegisterFunc      func(*grpc.Server)
	Logger            *zap.Logger
}

func RunGRPCServer(c* GRPCConfig) error {
	nameField := zap.String("name", c.Name)
	listener, err := net.Listen("tcp", c.Addr)
	if err != nil {
		c.Logger.Fatal("无法启动监听，错误=", nameField, zap.Error(err))
	}
	server := grpc.NewServer()
	c.RegisterFunc(server)
	c.Logger.Info("启动服务", nameField, zap.String("addr", c.Addr))
	return server.Serve(listener)
}
