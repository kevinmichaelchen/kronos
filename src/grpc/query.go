package grpc

import (
	"context"
	proto "github.com/IrisVR/kronos/pb"
)

func (s *Server) GetNumberOfLogins(ctx context.Context, in *proto.UserQuery) (*proto.CountResponse, error) {
	return &proto.CountResponse{
		Count: 0,
	}, nil
}

func (s *Server) GetUserSessionDuration(ctx context.Context, in *proto.UserQuery) (*proto.DurationResponse, error) {
	return &proto.DurationResponse{
		DurationMs: 0,
	}, nil
}
