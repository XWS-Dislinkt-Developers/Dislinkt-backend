package persistence

import (
	"errors"
	ftm "fmt"

	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/domain"
	"gorm.io/gorm"
)

type UserPostgresStore struct {
	db *gorm.DB
}

func NewUserPostgresStore(db *gorm.DB) (domain.UserStore, error) {
	err := db.AutoMigrate(&domain.User{})
	if err != nil {
		return nil, err
	}
	return &UserPostgresStore{
		db: db,
	}, nil
}

func (store *UserPostgresStore) Get(id string) (*domain.User, error) {

	var foundUser *domain.User

	store.db.Where("ID = ?", id).First(foundUser)

	if foundUser == nil {
		ftm.Println("[UserPostgresStore]:cant find user")
		return nil, errors.New("cant find user")
	}

	return foundUser, nil
}

func (store *UserPostgresStore) GetByUsername(username string) (*domain.User, error) {
	var foundUser *domain.User

	store.db.Where("username = ?", username).First(foundUser)

	if foundUser == nil {
		ftm.Println("[UserPostgresStore]:cant find user with username " + username)
		return nil, errors.New("cant find user")
	}

	return foundUser, nil
}

func (store *UserPostgresStore) Insert(user *domain.User) error {

	user, err := store.GetByUsername(user.Username)
	if err == nil {
		ftm.Println("[UserPostgresStore]:User is already registered: " + user.Username)
		return errors.New("user is already registered")
	}

	result := store.db.Create(user)
	if result.Error != nil {
		ftm.Println("[UserPostgresStore]Insert:cant insert user")
		return errors.New("cant insert user")
	}

	return nil
}

func (store *UserPostgresStore) GetAll() (*[]domain.User, error) {
	var users []domain.User
	result := store.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return &users, nil
}

func (store *UserPostgresStore) DeleteAll() {
	store.db.Session(&gorm.Session{AllowGlobalUpdate: true}).
		Delete(&domain.User{})
}
