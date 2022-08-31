package application

import (
	"errors"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_service/domain"
	logg "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_service/logger"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	store       domain.UserStore
	loggerInfo  *logg.Logger
	loggerError *logg.Logger
}

func NewUserService(store domain.UserStore, loggerInfo *logg.Logger, loggerError *logg.Logger) *UserService {
	return &UserService{
		store:       store,
		loggerInfo:  loggerInfo,
		loggerError: loggerError,
	}
}

func (service *UserService) Create(user *domain.User) error {
	err := service.store.Insert(user)
	if err != nil {
		service.loggerError.Logger.Error("User_service: CNSU ")
		println("Error in create method")
	}
	return err
}

func (service *UserService) Get(id int) (*domain.User, error) {
	return service.store.Get(id)
}

func (service *UserService) GetById(id int) (*domain.User, error) {
	return service.store.GetById(id)
}

func (service *UserService) IsProfilePrivate(id int) bool {
	user, _ := service.store.GetById(id)
	return user.IsPrivateProfile
}

func (service *UserService) GetByUsername(username string) (*domain.User, error) {
	return service.store.GetByUsername(username)
}

func (service *UserService) UpdateUser(dto domain.UpdateUserDto, currentUsername string, userID int) (*domain.User, error) {

	if currentUsername == dto.Username {
		println("Nije izmenio username")
		return service.store.UpdateUser(dto, userID)
	} else {
		println("Izmenio username")
		foundUserWithNewUsername, _ := service.GetByUsername(dto.Username)
		if foundUserWithNewUsername != nil {
			println("Izmenio username, ali korisnicko ime je vec zauzeto")
			return nil, errors.New("user with this username already exists")
		}
		println("Izmenio username, trebalo bi da sve je u redu")
		return service.store.UpdateUser(dto, userID)
	}

	/*


		if foundUserWithNewUsername != nil {
			if currentUsername == dto.Username {
				println("Nije izmenio")

				// nije izmenio username
				return service.store.UpdateUser(dto, userID)
			}else{
				return nil, errors.New("user with this username already exists")
			}
		}else{
			if currentUsername == dto.Username {
				// nije izmenio username
				return service.store.UpdateUser(dto, userID)
			}else{
				println("")
			}
		}










		if foundUserWithNewUsername != nil && foundUserWithNewUsername.UserId == userID {
			return service.store.UpdateUser(dto, userID)
		}
		service.loggerError.Logger.Error("User_service: USNAT  | UI " + strconv.Itoa(foundUser.ID))
		return nil, errors.New("Username is already taken")
	*/

}

func (service *UserService) UpdateUserWAE(dto domain.UpdateUserWAEDto, userID int) (*domain.User, error) {
	return service.store.UpdateUserWorkAndEducation(dto, userID)
}

func (service *UserService) UpdateUserSAI(dto domain.UpdateUserSAIDto, userID int) (*domain.User, error) {
	return service.store.UpdateUserSkillsAndInterests(dto, userID)
}

func (service *UserService) GetAll() (*[]domain.User, error) {
	return service.store.GetAll()
}

func (service *UserService) DeleteAll() {
	service.store.DeleteAll()
}

func (service *UserService) DeleteUser(id int) {
	service.store.DeleteUser(id)
}

func (service *UserService) ConfirmAccount(email string) {
	service.store.ConfirmAccount(email)
}

func (service *UserService) ChangePassword(email, password string) {
	service.store.ChangePassword(email, service.HashPassword(password))
}

func (service *UserService) HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 5)
	return string(bytes)
}

func (service *UserService) GetAllPublicProfiles() (*[]int, error) {
	users := make([]int, 0)
	allUsers, _ := service.store.GetAll()

	for _, user := range *allUsers {
		if user.IsPrivateProfile != true && user.IsItConfirmed == true {
			users = append(users, user.UserId)
		}
	}

	return &users, nil
}
