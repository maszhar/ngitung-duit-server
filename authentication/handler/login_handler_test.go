package handler_test

import (
	"context"
	"testing"
	"time"

	"github.com/djeniusinvfest/inventora/auth/entity"
	"github.com/djeniusinvfest/inventora/auth/handler"
	pb "github.com/djeniusinvfest/inventora/auth/proto"
	"github.com/djeniusinvfest/inventora/auth/repository"
	"github.com/djeniusinvfest/inventora/auth/util"
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

var now = time.Now()
var fakeRegisteredUser = &entity.User{
	Id:          "639919c5e753aedb94b7205e",
	Firstname:   "Ali",
	Lastname:    "Barbara",
	Email:       "mail@example.com",
	Password:    "$tTRiuH3tHoviChsig3mZjQ==$6D0Q2jAAPwK/yPM8YRWZV5EM9eg0pLunYRI4ZBQmHJ4=",
	ActivatedAt: &now,
}

func TestLoginInvalidRequest(t *testing.T) {
	p := generateLoginRequest()
	p.Email = ""

	ctrl, m := before(t)
	defer ctrl.Finish()

	handler := handler.NewHandler("", m)
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

	handler := handler.NewHandler("", m)
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

	handler := handler.NewHandler("", m)
	res, _ := handler.Login(context.Background(), p)

	if res.Result != pb.LoginResult_LOGIN_UNVERIFIED {
		t.Fatalf("Login(invalid message) = %v, wants %v", res.Result, pb.LoginResult_LOGIN_UNVERIFIED)
	}
}

func TestLoginSuccess(t *testing.T) {
	p := generateLoginRequest()
	jwtKey := "secret890189"

	ctrl, m := before(t)
	defer ctrl.Finish()

	m.EXPECT().
		Login(
			gomock.Any(),
			gomock.Any(),
		).
		Return(fakeRegisteredUser, nil)

	handler := handler.NewHandler(jwtKey, m)
	res, _ := handler.Login(context.Background(), p)

	if res.Result != pb.LoginResult_LOGIN_SUCCESS {
		t.Fatalf("Login(invalid message) = %v, wants %v", res.Result, pb.LoginResult_LOGIN_SUCCESS)
	}

	_, err := util.ParseAccessToken(jwtKey, res.AccessToken)
	if err != nil {
		t.Fatal(err)
	}
}
