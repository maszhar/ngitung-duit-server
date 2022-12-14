package repository

import (
	"errors"

	"github.com/djeniusinvfest/inventora/auth/entity"
	"github.com/djeniusinvfest/inventora/auth/model"
	"github.com/djeniusinvfest/inventora/auth/util"
	"go.mongodb.org/mongo-driver/bson"
)

type AuthRepository interface {
	RegisterUser(e *entity.User) error
	Login(email string, password string) (*entity.User, error)
}

type authRepository struct {
	userModel model.UserModel
}

var ErrEmailConflict = errors.New(("auth repo: email is used by another user"))

func (r *authRepository) RegisterUser(user *entity.User) error {
	// Check existing email
	filter := bson.D{{"email", user.Email}}
	foundEmail, err := r.userModel.FindOne(filter, true)
	if err != nil {
		return err
	}
	if foundEmail != nil {
		return ErrEmailConflict
	}

	// Hash password
	digest, err := util.HashString(user.Password)
	if err != nil {
		return err
	}
	user.Password = digest

	err = r.userModel.Create(user)

	return err
}

var ErrInvalidCreds = errors.New("invalid credentials")
var ErrUnverifiedAccount = errors.New("unverified account")

func (r *authRepository) Login(email string, password string) (*entity.User, error) {
	// fetch user data
	filter := bson.D{{"email", email}}
	user, err := r.userModel.FindOne(filter, false)
	if err != nil {
		return nil, err
	}

	// wrong email
	if user == nil {
		return nil, ErrInvalidCreds
	}

	// wrong password
	if !util.VerifyDigest(password, user.Password) {
		return nil, ErrInvalidCreds
	}

	// account unverified
	if user.ActivatedAt == nil {
		return nil, ErrUnverifiedAccount
	}

	return user, nil
}

func NewAuthRepo(um model.UserModel) AuthRepository {
	return &authRepository{
		userModel: um,
	}
}
