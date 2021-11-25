package auth

import authpb "service/auth/api/gen/v1"

type Service struct {
	*authpb.UnimplementedAuthServiceServer
}

