package repository_test

import (
	"testing"
	"time"

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

var now = time.Now()
var fakeRegisteredUser = &entity.User{
	Firstname:   "Ali",
	Lastname:    "Barbara",
	Email:       "mail@example.com",
	Password:    "$tTRiuH3tHoviChsig3mZjQ==$6D0Q2jAAPwK/yPM8YRWZV5EM9eg0pLunYRI4ZBQmHJ4=",
	ActivatedAt: &now,
}

func TestRegisterUser(t *testing.T) {
	e := generateUserEntity()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_model.NewMockUserModel(ctrl)
	m.EXPECT().
		FindOne(
			gomock.Any(),
			gomock.Eq(true),
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
		FindOne(
			gomock.Any(),
			gomock.Eq(true),
		).
		Return(&entity.User{}, nil)

	authRepo := repository.NewAuthRepo(m)
	err := authRepo.RegisterUser(e)
	if err != repository.ErrEmailConflict {
		t.Fatalf("RegisterUser(exists email) = %v, wants %v", err, repository.ErrEmailConflict)
	}
}

func TestLoginSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_model.NewMockUserModel(ctrl)
	m.EXPECT().
		FindOne(
			gomock.Any(),
			gomock.Eq(false),
		).
		Return(fakeRegisteredUser, nil)

	authRepo := repository.NewAuthRepo(m)
	user, _ := authRepo.Login("mail@example.com", "password")
	if user == nil {
		t.Fatalf("Login(valid) = %v, wants %v", user, fakeRegisteredUser)
	}
}

func TestLoginWrongEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_model.NewMockUserModel(ctrl)
	m.EXPECT().
		FindOne(
			gomock.Any(),
			gomock.Eq(false),
		).
		Return(nil, nil)

	authRepo := repository.NewAuthRepo(m)
	_, err := authRepo.Login("mail1@example.com", "password")
	if err != repository.ErrInvalidCreds {
		t.Fatalf("Login(valid) = %v, wants %v", err, repository.ErrInvalidCreds)
	}
}

func TestLoginWrongPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_model.NewMockUserModel(ctrl)
	m.EXPECT().
		FindOne(
			gomock.Any(),
			gomock.Eq(false),
		).
		Return(fakeRegisteredUser, nil)

	authRepo := repository.NewAuthRepo(m)
	_, err := authRepo.Login("mail@example.com", "password12")
	if err != repository.ErrInvalidCreds {
		t.Fatalf("Login(valid) = %v, wants %v", err, repository.ErrInvalidCreds)
	}
}

func TestLoginUnverified(t *testing.T) {
	fakeUser := fakeRegisteredUser
	fakeUser.ActivatedAt = nil

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_model.NewMockUserModel(ctrl)
	m.EXPECT().
		FindOne(
			gomock.Any(),
			gomock.Eq(false),
		).
		Return(fakeUser, nil)

	authRepo := repository.NewAuthRepo(m)
	_, err := authRepo.Login("mail@example.com", "password")
	if err != repository.ErrUnverifiedAccount {
		t.Fatalf("Login(valid) = %v, wants %v", err, repository.ErrUnverifiedAccount)
	}
}
