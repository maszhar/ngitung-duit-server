package handler

import (
	pb "github.com/djeniusinvfest/inventora/auth/proto"
	"github.com/djeniusinvfest/inventora/auth/repository"
)

type Handler struct {
	pb.UnimplementedAuthenticationServer

	authRepo repository.AuthRepository
	jwtKey   string
}

func NewHandler(jwtKey string, authRepo repository.AuthRepository) *Handler {
	h := &Handler{
		authRepo: authRepo,
		jwtKey:   jwtKey,
	}
	return h
}
