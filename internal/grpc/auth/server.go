package auth

import (
	"context"
	auth1 "github.com/mursalbekov1/proto/gen/go/auth"
	"google.golang.org/grpc"
)

type ServerAPI struct {
	auth1.UnimplementedAuthServer
}

func Register(gRPC *grpc.Server) {
	auth1.RegisterAuthServer(gRPC, &ServerAPI{})
}

func (s *ServerAPI) Login(
	ctx context.Context,
	req *auth1.LoginRequest,
) (*auth1.LoginResponse, error) {
	panic("implement me")
}

func (s *ServerAPI) Register(
	ctx context.Context,
	req *auth1.RegisterRequest,
) (*auth1.RegisterResponse, error) {
	panic("implement me")
}

func (s *ServerAPI) IsAdmin(
	ctx context.Context,
	req *auth1.IsAdminRequest,
) (*auth1.IsAdminResponse, error) {
	panic("implement me")
}
