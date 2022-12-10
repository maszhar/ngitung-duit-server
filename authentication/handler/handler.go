package handler

import (
	pb "github.com/djeniusinvfest/inventora/auth/proto"
)

type Handler struct {
	pb.UnimplementedAuthenticationServer
}

func NewHandler() *Handler {
	h := &Handler{}
	return h
}
