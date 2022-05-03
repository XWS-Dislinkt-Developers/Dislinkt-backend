package application

import (
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_post_servis/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserPostService struct {
	store domain.UserPostStore
}

func NewUserPostService(store domain.UserPostStore) *UserPostService {
	return &UserPostService{
		store: store,
	}
}

func (service *UserPostService) Get(id primitive.ObjectID) (*domain.UserPost, error) {
	return service.store.Get(id)
}

func (service *UserPostService) GetAll() ([]*domain.UserPost, error) {
	return service.store.GetAll()
}

func (service *UserPostService) Create(userPost *domain.UserPost) error {
	userPost.CreatedAt = time.Now()
	err := service.store.Insert(userPost)
	if err != nil {
		return err
	}
	return nil
}
