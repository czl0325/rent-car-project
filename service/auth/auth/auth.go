package auth

import (
	"context"
	authpb "service/auth/api/gen/v1"
	"service/auth/dao"
)

type Service struct {
	*authpb.UnimplementedAuthServiceServer
	Mongo *dao.Mongo
}

func (s *Service) Login(c context.Context, request *authpb.LoginRequest) (*authpb.UserInfo, error) {
	return s.Mongo.LoginWithRegister(c, request)
}
