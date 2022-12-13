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
}

type authRepository struct {
	userModel model.UserModel
}

var ErrEmailConflict = errors.New(("auth repo: email is used by another user"))

func (r *authRepository) RegisterUser(user *entity.User) error {

	foundEmail, err := r.FindUserByEmail(user.Email)
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

func (r *authRepository) FindUserByEmail(email string) (*entity.User, error) {
	filter := bson.D{{"email", email}}

	user, err := r.userModel.FindOneWithDeleted(filter)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func NewAuthRepo(um model.UserModel) AuthRepository {
	return &authRepository{
		userModel: um,
	}
}
