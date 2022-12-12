package handler

import (
	"context"

	pb "github.com/djeniusinvfest/inventora/auth/proto"
	"github.com/djeniusinvfest/inventora/auth/repository"
)

func (h *Handler) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	err := h.authRepo.RegisterUser(in)
	if err != nil {
		if err == repository.ErrEmailConflict {
			return &pb.RegisterResponse{
				Result:  pb.Result_INVALID_FIELDS,
				Message: err.Error(),
			}, nil
		}
		return nil, err
	}

	return &pb.RegisterResponse{
		Result:  pb.Result_SUCCESS,
		Message: "",
	}, nil
}
