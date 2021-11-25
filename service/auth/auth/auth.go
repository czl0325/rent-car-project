package auth

import (
	"context"
	authpb "service/auth/api/gen/v1"
)

type Service struct {
	*authpb.UnimplementedAuthServiceServer
}

func (s *Service) Login(c context.Context, request *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	return &authpb.LoginResponse{
		Token:    request.Code,
		ExpireIn: 1000,
	}, nil
}
