package repository_test

import (
	"testing"

	"github.com/djeniusinvfest/inventora/auth/entity"
	"github.com/djeniusinvfest/inventora/auth/model/mock_model"
	"github.com/djeniusinvfest/inventora/auth/repository"
	"github.com/golang/mock/gomock"
	"github.com/jaswdr/faker"
)

func generateUserEntity() *entity.User {
	faker := faker.New()
	return &entity.User{
		Firstname: faker.Person().FirstName(),
		Lastname:  faker.Person().LastName(),
		Email:     faker.Internet().Email(),
		Password:  faker.Internet().Password(),
	}
}

func TestRegisterUser(t *testing.T) {
	e := generateUserEntity()

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
	err := authRepo.RegisterUser(e)
	if err != nil {
		t.Fatalf("RegisterUser(valid) = %v, wants no error", err)
	}
}

func TestRegisterUserEmailExists(t *testing.T) {
	e := generateUserEntity()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_model.NewMockUserModel(ctrl)
	m.EXPECT().
		FindOneWithDeleted(
			gomock.Any(),
		).
		Return(&entity.User{}, nil)

	authRepo := repository.NewAuthRepo(m)
	err := authRepo.RegisterUser(e)
	if err != repository.ErrEmailConflict {
		t.Fatalf("RegisterUser(exists email) = %v, wants %v", err, repository.ErrEmailConflict)
	}
}
