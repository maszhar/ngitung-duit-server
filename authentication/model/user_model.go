package model

import (
	"context"
	"log"
	"time"

	"github.com/djeniusinvfest/inventora/auth/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserModel interface {
	Create(user *entity.User) error
	FindOneWithDeleted(filter interface{}) (*entity.User, error)
}

type userModel struct {
	coll *mongo.Collection
}

func (um *userModel) Create(user *entity.User) error {
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now
	result, err := um.coll.InsertOne(
		context.Background(),
		user,
	)
	if err != nil {
		return err
	}

	newId, err := parseStringId(result.InsertedID)
	if err != nil {
		return err
	}

	user.Id = newId
	return nil
}

func (um *userModel) FindOneWithDeleted(filter interface{}) (*entity.User, error) {
	var result entity.User
	err := um.coll.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}

func (um *userModel) defineIndexes() {
	// email unique
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{"email", 1}},
		Options: options.Index().SetUnique(true),
	}
	_, err := um.coll.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		log.Println("UserModel: failed to create email index")
	}
}

func NewUserModel(db *mongo.Database) UserModel {
	user := &userModel{
		coll: db.Collection("user"),
	}

	user.defineIndexes()
	return user
}
