package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserPostStore interface {
	Get(id primitive.ObjectID) (*UserPost, error)
	GetAll() ([]*UserPost, error)
	Insert(userPost *UserPost) error
	DeleteAll()
}
