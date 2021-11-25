package server

import "go.uber.org/zap"

func NewZapLogger() (*zap.Logger, error)  {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.TimeKey = ""
	return config.Build()
}
