package repository_test

import (
	"testing"

	"github.com/djeniusinvfest/inventora/auth/entity"
	"github.com/djeniusinvfest/inventora/auth/model/mock_model"
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

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_model.NewMockUserModel(ctrl)
	m.EXPECT().
		FindOneWithDeleted(
			gomock.Any(),
		).
		Return(nil, nil)
	m.EXPECT().
		Create(
			gomock.Any(),
		).
		Return(nil)

	authRepo := repository.NewAuthRepo(m)
	err := authRepo.RegisterUser(p)
	if err != nil {
		t.Fatalf("RegisterUser(valid) = %v, wants no error", err)
	}
}

func TestRegisterUserEmailExists(t *testing.T) {
	p := generateRegisterRequest()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_model.NewMockUserModel(ctrl)
	m.EXPECT().
		FindOneWithDeleted(
			gomock.Any(),
		).
		Return(&entity.User{}, nil)

	authRepo := repository.NewAuthRepo(m)
	err := authRepo.RegisterUser(p)
	if err != repository.ErrEmailConflict {
		t.Fatalf("RegisterUser(exists email) = %v, wants %v", err, repository.ErrEmailConflict)
	}
}
