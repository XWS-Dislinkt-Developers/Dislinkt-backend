package persistence

import (
	"errors"
	ftm "fmt"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/domain"
	logg "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/logger"
	"gorm.io/gorm"
	"strconv"
)

type UserPostgresStore struct {
	db          *gorm.DB
	loggerInfo  *logg.Logger
	loggerError *logg.Logger
}

func NewUserPostgresStore(db *gorm.DB, loggerInfo *logg.Logger, loggerError *logg.Logger) (domain.UserStore, error) {
	err := db.AutoMigrate(&domain.User{})
	if err != nil {
		return nil, err
	}
	return &UserPostgresStore{db: db,
		loggerInfo:  loggerInfo,
		loggerError: loggerError}, nil
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

func (store *UserPostgresStore) GetById(id int) (*domain.User, error) {
	var foundUser *domain.User
	users, err := store.GetAll()
	if err != nil {
		store.loggerError.Logger.Errorf("User_postgres_store: GetById - failed method - can't find user with user id " + strconv.Itoa(id))

		return nil, errors.New("[UserPostgresStore-GetByUsername(username)]: There's no user.")
	}
	for _, user := range *users {
		if user.ID == id {
			return &user, nil
		}
	}
	if foundUser == nil {
		store.loggerError.Logger.Errorf("User_postgres_store: GetById - failed method - can't find user with user id " + strconv.Itoa(id))
		ftm.Println("[UserPostgresStore-GetById(id)]: Can't find user with this id: " + string(id))
		return nil, errors.New("ERR - [UserPostgresStore-GetById(id)]: Can't find user with this id: " + string(id))
	}
	store.loggerInfo.Logger.Infof("User_postgres_store: GetById - User id " + strconv.Itoa(id))

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
		store.loggerError.Logger.Errorf("User_postgres_store: GetByUsername - failed method - no users")

		return nil, errors.New("[UserPostgresStore-GetByUsername(username)]: There's no user.")
	}
	for _, user := range *users {
		if user.Username == username {
			return &user, nil
		}
	}
	if foundUser == nil {
		store.loggerError.Logger.Errorf("User_postgres_store: GetByUsername - failed method - Can't find user with this username ")

		ftm.Println("[UserPostgresStore-GetByUsername(username)]: Can't find user with this username ")
		return nil, errors.New("ERR - [UserPostgresStore-GetByUsername(username)]: Can't find user with sent username ")
	}
	store.loggerInfo.Logger.Infof("User_postgres_store: GetByUsername - user with user id " + strconv.Itoa(foundUser.ID) + " is found.")

	return foundUser, nil
}

func (store *UserPostgresStore) GetByEmail(email string) (*domain.User, error) {
	var foundUser *domain.User
	users, err := store.GetAll()
	if err != nil {
		return nil, errors.New("[UserPostgresStore-GetByUsername(username)]: There's no user.")
	}
	for _, user := range *users {
		if user.Email == email {
			return &user, nil
		}
	}
	if foundUser == nil {
		store.loggerError.Logger.Errorf("User_postgres_store: GetByEmail - failed method - Can't find user with this email ")

		ftm.Println("[UserPostgresStore-GetByEmail(email)]: Can't find user with this email: " + email)
		return nil, errors.New("ERR - [UserPostgresStore-GetByEmail(email)]: Can't find user with this email: " + email)
	}
	store.loggerInfo.Logger.Infof("User_postgres_store: GetByEmail - user with user id " + strconv.Itoa(foundUser.ID) + " is found.")

	return foundUser, nil
}

func (store *UserPostgresStore) UpdateUser(dto domain.UpdateUserDto, userID int) (*domain.User, error) {
	var user domain.User
	user.ID = userID
	store.db.First(&user)
	user.Username = dto.Username
	user.Name = dto.Name
	user.Gender = dto.Gender
	user.Email = dto.Email
	user.PhoneNumber = dto.PhoneNumber
	user.Biography = dto.Biography
	user.DateOfBirth = dto.DateOfBirth
	store.db.Save(&user)
	store.loggerInfo.Logger.Infof("User_postgres_store: UpdateUser - User id " + strconv.Itoa(userID) + " changed data in database!")

	return nil, nil
}

func (store *UserPostgresStore) ConfirmAccount(idUser int) (*domain.User, error) {
	var user domain.User
	user.ID = idUser
	store.db.First(&user)
	user.IsItConfirmed = true
	store.db.Save(&user)
	store.loggerInfo.Logger.Infof("User_postgres_store: ConfirmAccount - User id " + strconv.Itoa(idUser) + " confirmed account")

	return nil, nil
}

func (store *UserPostgresStore) UpdateUserWorkAndEducation(dto domain.UpdateUserWAEDto, userId int) (*domain.User, error) {
	var user domain.User
	user.ID = userId
	store.db.First(&user)
	user.Work = dto.Work
	user.Education = dto.Education
	store.db.Save(&user)
	store.loggerInfo.Logger.Infof("User_postgres_store: UpdateUserWorkAndEducation - User id " + strconv.Itoa(userId) + " change data in database!")

	return nil, nil
}
func (store *UserPostgresStore) UpdateUserSkillsAndInterests(dto domain.UpdateUserSAIDto, userId int) (*domain.User, error) {
	var user domain.User
	user.ID = userId
	store.db.First(&user)
	user.Skills = dto.Skills
	user.Interests = dto.Interests
	store.db.Save(&user)
	store.loggerInfo.Logger.Infof("User_postgres_store: UpdateUserSkillsAndInterests - User id " + strconv.Itoa(userId) + " change data in database!")

	return nil, nil
}

func (store *UserPostgresStore) UpdatePassword(userId int, password string) {
	var user domain.User
	user.ID = userId
	store.db.First(&user)
	user.Password = password
	store.db.Save(&user)
	store.loggerInfo.Logger.Infof("User_postgres_store: UpdatePassword - User id " + strconv.Itoa(userId) + " changed password! ")

}

func (store *UserPostgresStore) Insert(user *domain.User) error {
	_, err := store.GetByUsername(user.Username)
	if err == nil {
		store.loggerError.Logger.Errorf("User_postgres_store: Insert - failed method - user " + user.Username + " is already registred! ")

		ftm.Println("[UserPostgresStore-Insert(user)]: User is already registered: " + user.Username)
		return errors.New("ERR - [UserPostgresStore-Insert(user)]: User is already registered: " + user.Username)
	}

	result := store.db.Create(user)
	if result.Error != nil {
		store.loggerError.Logger.Errorf("User_postgres_store: Insert - failed method - user " + user.Username + " can't be saved!")

		ftm.Println("[UserPostgresStore-Insert(user)]: Can't insert user.")
		return errors.New("ERR - [UserPostgresStore-Insert(user)]: Can't insert user. ")
	}
	store.loggerInfo.Logger.Infof("User_postgres_store: Insert - User " + user.Username + " is registred! ")

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
