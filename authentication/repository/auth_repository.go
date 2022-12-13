package repository

import (
	"errors"

	"github.com/djeniusinvfest/inventora/auth/entity"
	"github.com/djeniusinvfest/inventora/auth/model"
	pb "github.com/djeniusinvfest/inventora/auth/proto"
	"go.mongodb.org/mongo-driver/bson"
)

type AuthRepository interface {
	RegisterUser(p *pb.RegisterRequest) error
}

type authRepository struct {
	userModel model.UserModel
}

var ErrEmailConflict = errors.New(("auth repo: email is used by another user"))

func (r *authRepository) RegisterUser(p *pb.RegisterRequest) error {

	foundEmail, err := r.FindUserByEmail(p.Email)
	if err != nil {
		return err
	}
	if foundEmail != nil {
		return ErrEmailConflict
	}

	user := entity.User{
		Firstname: p.FirstName,
		Lastname:  p.LastName,
		Email:     p.Email,
		Password:  p.Password,
	}
	err = r.userModel.Create(&user)

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
