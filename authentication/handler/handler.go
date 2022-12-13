package handler

import (
	pb "github.com/djeniusinvfest/inventora/auth/proto"
	"github.com/djeniusinvfest/inventora/auth/repository"
)

type Handler struct {
	pb.UnimplementedAuthenticationServer

	authRepo repository.AuthRepository
}

func NewHandler(authRepo repository.AuthRepository) *Handler {
	h := &Handler{
		authRepo: authRepo,
	}
	return h
}
