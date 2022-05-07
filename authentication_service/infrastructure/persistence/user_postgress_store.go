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
	return &UserPostgresStore{db: db}, nil
}

func (store *UserPostgresStore) Get(id int) (*domain.User, error) { // TODO: why not integer (id)??
	var foundUser *domain.User
	store.db.Where("ID = ?", id).First(foundUser)
	if foundUser == nil {
		ftm.Println("[UserPostgresStore-Get(id)]: Can't find user")
		return nil, errors.New("ERR-[UserPostgresStore-Get(id)]: Can't find user ")
	}
	return foundUser, nil
}

func (store *UserPostgresStore) GetByUsername(username string) (*domain.User, error) {
	var foundUser *domain.User
	//result := store.db.Where("username = ?", username).First(foundUser)
	// ftm.Println("[UserPostgresStore]:username ", username)
	// ftm.Println("[UserPostgresStore]:result found ", result.RowsAffected)
	//var users []domain.User
	users, err := store.GetAll()
	if err != nil {
		return nil, errors.New("[UserPostgresStore-GetByUsername(username)]: There's no user.")
	}
	for _, user := range *users {
		if user.Username == username {
			return &user, nil
		}
	}
	if foundUser == nil {
		ftm.Println("[UserPostgresStore-GetByUsername(username)]: Can't find user with this username: " + username)
		return nil, errors.New("ERR - [UserPostgresStore-GetByUsername(username)]: Can't find user with this username: " + username)
	}
	return foundUser, nil
}

func (store *UserPostgresStore) UpdateUser(dto domain.UpdateUserDto) (*domain.User, error) {
	var user *domain.User
	store.db.First(&user)
	user.Username = dto.Username
	user.Name = dto.Name
	user.Gender = dto.Gender
	user.Email = dto.Email
	user.PhoneNumber = dto.PhoneNumber
	user.Biography = dto.Biography
	user.DateOfBirth = dto.DateOfBirth
	store.db.Save(&user)

	return user, nil
}

func (store *UserPostgresStore) UpdateUserWorkAndEducation(dto domain.UpdateUserWAEDto, userId int) (*domain.User, error) {
	var user domain.User
	user.ID = userId
	store.db.First(&user)
	user.Work = dto.Work
	user.Education = dto.Education
	store.db.Save(&user)
	return nil, nil
}
func (store *UserPostgresStore) UpdateUserSkillsAndInterests(dto domain.UpdateUserSAIDto, userId int) (*domain.User, error) {
	var user domain.User
	user.ID = userId
	store.db.First(&user)
	user.Skills = dto.Skills
	user.Interests = dto.Interests
	store.db.Save(&user)
	return nil, nil
}
func (store *UserPostgresStore) Insert(user *domain.User) error {
	_, err := store.GetByUsername(user.Username)
	if err == nil {
		ftm.Println("[UserPostgresStore-Insert(user)]: User is already registered: " + user.Username)
		return errors.New("ERR - [UserPostgresStore-Insert(user)]: User is already registered: " + user.Username)
	}
	result := store.db.Create(user)
	if result.Error != nil {
		ftm.Println("[UserPostgresStore-Insert(user)]: Can't insert user.")
		return errors.New("ERR - [UserPostgresStore-Insert(user)]: Can't insert user. ")
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
