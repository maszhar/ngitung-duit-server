package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/djeniusinvfest/inventora/auth/model"
	pb "github.com/djeniusinvfest/inventora/auth/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepository struct {
	db       *mongo.Database
	userColl *mongo.Collection
}

func NewAuthRepo(db *mongo.Database) *AuthRepository {
	return &AuthRepository{
		db:       db,
		userColl: db.Collection("user"),
	}
}

var ErrEmailConflict = errors.New(("auth_repo: email is used by another user"))

func (r *AuthRepository) RegisterUser(p *pb.RegisterRequest) error {
	_, err := r.FindUserByEmail(p.Email)
	if err == nil {
		return ErrEmailConflict
	} else {
		if err != mongo.ErrNoDocuments {
			return err
		}
	}

	now := time.Now()
	user := model.User{
		Firstname: p.FirstName,
		Lastname:  p.LastName,
		Email:     p.Email,
		Password:  p.Password,
		CreatedAt: now,
		UpdatedAt: now,
	}
	r.userColl.InsertOne(
		context.Background(),
		user,
	)

	return nil
}

func (r *AuthRepository) FindUserByEmail(email string) (*model.User, error) {
	filter := bson.D{{"email", email}}

	var user model.User
	err := r.userColl.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		fmt.Printf("%v", err)
		return nil, err
	}

	return &user, nil
}
