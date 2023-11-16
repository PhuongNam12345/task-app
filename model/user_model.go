package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Fullname string             `json:"fullname,omitempty" validate:"required"`
	Email    string             `json:"email,omitempty" validate:"required"`
	Phone    string             `json:"phone,omitempty" validate:"required"`
	Address  string             `json:"address,omitempty" validate:"required"`
	Roles    string             `json:"roles" validate:"required"`
}
