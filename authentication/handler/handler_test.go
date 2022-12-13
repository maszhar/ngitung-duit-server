package handler_test

import (
	"testing"

	"github.com/djeniusinvfest/inventora/auth/repository/mock_repository"
	"github.com/golang/mock/gomock"
)

func before(t *testing.T) (*gomock.Controller, *mock_repository.MockAuthRepository) {
	ctrl := gomock.NewController(t)
	m := mock_repository.NewMockAuthRepository(ctrl)
	return ctrl, m
}
