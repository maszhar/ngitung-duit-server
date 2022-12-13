package model

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ErrInvalidDocId = errors.New("model: invalid document id")

func parseStringId(in interface{}) (string, error) {
	if id, ok := in.(primitive.ObjectID); ok {
		return id.String(), nil
	} else {
		return "", ErrInvalidDocId
	}
}
