package handler

import (
	"context"

	pb "github.com/djeniusinvfest/inventora/auth/proto"
)

func (h *Handler) Hello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "authentication server is running",
	}, nil
}
