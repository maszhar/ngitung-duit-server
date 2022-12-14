package entity

import (
	"time"
)

type User struct {
	Id          string     `bson:"_id,omitempty"`
	Username    *string    `bson:"username"`
	Firstname   string     `bson:"first_name"`
	Lastname    string     `bson:"last_name"`
	Email       string     `bson:"email"`
	Password    string     `bson:"password"`
	CreatedAt   time.Time  `bson:"created_at"`
	UpdatedAt   time.Time  `bson:"updated_at"`
	ActivatedAt *time.Time `bson:"activated_at"`
	DeletedAt   *time.Time `bson:"deleted_at"`
}
