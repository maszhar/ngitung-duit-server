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

func generateRegisterRequest() *pb.RegisterRequest {
	faker := faker.New()
	return &pb.RegisterRequest{
		FirstName: faker.Person().FirstName(),
		LastName:  faker.Person().LastName(),
		Email:     faker.Internet().Email(),
		Password:  faker.Internet().Password(),
		AgreeTos:  true,
	}
}

func TestRegisterUser(t *testing.T) {
	p := generateRegisterRequest()

	ctrl, m := before(t)
	defer ctrl.Finish()

	m.
		EXPECT().
		RegisterUser(gomock.Any()).
		Return(nil)

	handler := handler.NewHandler("", m)
	res, _ := handler.Register(context.Background(), p)

	if res.Result != pb.Result_SUCCESS {
		t.Fatalf("RegisterUser(valid message) = %v, wants %v", res.Result, pb.Result_SUCCESS)
	}
}

func TestRegisterUserEmailExists(t *testing.T) {
	p := generateRegisterRequest()

	ctrl, m := before(t)
	defer ctrl.Finish()

	m.
		EXPECT().
		RegisterUser(gomock.Any()).
		Return(repository.ErrEmailConflict)

	handler := handler.NewHandler("", m)
	res, _ := handler.Register(context.Background(), p)

	if res.Result != pb.Result_INVALID_FIELDS {
		t.Fatalf("RegisterUser(exists email) = %v, wants %v", res.Result, pb.Result_INVALID_FIELDS)
	}
}

func TestRegisterUserInvalidParams(t *testing.T) {
	p := generateRegisterRequest()
	p.FirstName = ""

	ctrl, m := before(t)
	defer ctrl.Finish()

	handler := handler.NewHandler("", m)
	res, _ := handler.Register(context.Background(), p)

	if res.Result != pb.Result_INVALID_FIELDS {
		t.Fatalf("RegisterUser(invalid message) = %v, wants %v", res.Result, pb.Result_INVALID_FIELDS)
	}
}
