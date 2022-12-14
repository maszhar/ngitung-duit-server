package handler

import (
	"context"

	"github.com/djeniusinvfest/inventora/auth/proto"
	"github.com/djeniusinvfest/inventora/auth/repository"
	"github.com/djeniusinvfest/inventora/auth/util"
	"github.com/djeniusinvfest/inventora/auth/validator"
)

func (h *Handler) Login(ctx context.Context, p *proto.LoginRequest) (*proto.LoginResponse, error) {
	// Validate request
	err := validator.ValidateLogin(p)
	if err != nil {
		return &proto.LoginResponse{
			Result:  proto.LoginResult_LOGIN_INVALID_FIELDS,
			Message: err.Error(),
		}, nil
	}

	// Login
	user, err := h.authRepo.Login(p.Email, p.Password)
	if err != nil {
		if err == repository.ErrInvalidCreds {
			return &proto.LoginResponse{
				Result: proto.LoginResult_LOGIN_INCORRECT_DATA,
			}, nil
		}
		if err == repository.ErrUnverifiedAccount {
			return &proto.LoginResponse{
				Result: proto.LoginResult_LOGIN_UNVERIFIED,
			}, nil
		}
		return nil, err
	}

	accessToken, err := util.CreateAccessToken(h.jwtKey, user.Id)
	if err != nil {
		return nil, err
	}

	return &proto.LoginResponse{
		Result:      proto.LoginResult_LOGIN_SUCCESS,
		AccessToken: accessToken,
	}, nil
}
