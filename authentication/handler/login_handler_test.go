package handler_test

import (
	"context"
	"testing"

	"github.com/djeniusinvfest/inventora/auth/handler"
	pb "github.com/djeniusinvfest/inventora/auth/proto"
	"github.com/djeniusinvfest/inventora/auth/repository"
	"github.com/golang/mock/gomock"
	"github.com/jaswdr/faker"
)

func generateLoginRequest() *pb.LoginRequest {
	faker := faker.New()
	return &pb.LoginRequest{
		Email:    faker.Internet().Email(),
		Password: faker.Internet().Password(),
	}
}

func TestLoginInvalidRequest(t *testing.T) {
	p := generateLoginRequest()
	p.Email = ""

	ctrl, m := before(t)
	defer ctrl.Finish()

	handler := handler.NewHandler(m)
	res, _ := handler.Login(context.Background(), p)

	if res.Result != pb.LoginResult_LOGIN_INVALID_FIELDS {
		t.Fatalf("Login(invalid message) = %v, wants %v", res.Result, pb.LoginResult_LOGIN_INVALID_FIELDS)
	}
}

func TestLoginInvalidCreds(t *testing.T) {
	p := generateLoginRequest()

	ctrl, m := before(t)
	defer ctrl.Finish()

	m.EXPECT().
		Login(
			gomock.Any(),
			gomock.Any(),
		).
		Return(nil, repository.ErrInvalidCreds)

	handler := handler.NewHandler(m)
	res, _ := handler.Login(context.Background(), p)

	if res.Result != pb.LoginResult_LOGIN_INCORRECT_DATA {
		t.Fatalf("Login(invalid message) = %v, wants %v", res.Result, pb.LoginResult_LOGIN_INCORRECT_DATA)
	}
}

func TestLoginUnferified(t *testing.T) {
	p := generateLoginRequest()

	ctrl, m := before(t)
	defer ctrl.Finish()

	m.EXPECT().
		Login(
			gomock.Any(),
			gomock.Any(),
		).
		Return(nil, repository.ErrUnverifiedAccount)

	handler := handler.NewHandler(m)
	res, _ := handler.Login(context.Background(), p)

	if res.Result != pb.LoginResult_LOGIN_UNVERIFIED {
		t.Fatalf("Login(invalid message) = %v, wants %v", res.Result, pb.LoginResult_LOGIN_UNVERIFIED)
	}
}
