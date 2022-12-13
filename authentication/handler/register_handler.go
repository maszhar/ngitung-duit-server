package handler

import (
	"context"

	"github.com/djeniusinvfest/inventora/auth/entity"
	pb "github.com/djeniusinvfest/inventora/auth/proto"
	"github.com/djeniusinvfest/inventora/auth/repository"
	"github.com/djeniusinvfest/inventora/auth/validator"
)

func (h *Handler) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	err := validator.ValidateRegister(in)
	if err != nil {
		return &pb.RegisterResponse{
			Result:  pb.Result_INVALID_FIELDS,
			Message: err.Error(),
		}, nil
	}

	user := entity.User{
		Firstname: in.FirstName,
		Lastname:  in.LastName,
		Email:     in.Email,
		Password:  in.Password,
	}

	err = h.authRepo.RegisterUser(&user)

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
