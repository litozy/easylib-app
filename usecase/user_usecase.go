package usecase

import (
	"easylib-go/model"
	"easylib-go/repository"
	"easylib-go/utils"
	"fmt"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	AddUser(*model.User) error
	GetUserById(string) (*model.User, error)
	UpdateUser(*model.User, *gin.Context) error
	DeleteUser(string) error
}

type userUsecase struct {
	usrRepo repository.UserRepository
}

func (usrUseCase *userUsecase) GetUserById(id string) (*model.User, error) {
	return usrUseCase.usrRepo.GetUserById(id)
}

func (usrUseCase *userUsecase) AddUser(usr *model.User) error {
	if usr.Username == "" {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: "Name cannot be empty",
		}
	}
	if len(usr.Username) < 3 || len(usr.Username) > 20 {
		return &utils.AppError{
			ErrorCode:    2,
			ErrorMessage: "Name must be between 3 and 20 characters",
		}
	}
	if usr.Password == "" {
		return &utils.AppError{
			ErrorCode:    3,
			ErrorMessage: "Password cannot be empty",
		}
	}
	user,_ := usrUseCase.usrRepo.GetUserByUsername(usr.Username)
	if user != nil {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: fmt.Sprintf("User data with the name %v already exists", usr.Username),
		}
	}
	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
    }
	CreatedAt := time.Now().UTC()
	UpdatedAt := time.Now().UTC()

	usr.Id = uuid.New().String()
	usr.CreatedAt = CreatedAt.Format("2006-01-02 15:04:05")
	usr.UpdatedAt = UpdatedAt.Format("2006-01-02 15:04:05")
	usr.Password = string(hashedPassword)
   return usrUseCase.usrRepo.AddUser(usr)
}

func (usrUseCase *userUsecase) UpdateUser(usr *model.User, ctx *gin.Context) error {
	session := sessions.Default(ctx)
	existSession := session.Get("Id")
	usr.Id = existSession.(string)
	if usr.Username == "" {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: "Name cannot be empty",
		}
	}
	if usr.Password == "" {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: "Password cannot be empty",
		}
	}

	existDataUsr, _ := usrUseCase.usrRepo.GetUserByUsername(usr.Username)
	if existDataUsr != nil {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: fmt.Sprintf("User data with the username %v already exists", usr.Username),
		}
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("userUsecase.GenerateFromPassword(): %w", err)
	}
	UpdatedAt := time.Now().UTC()
	usr.UpdatedAt = UpdatedAt.Format("2006-01-02 15:04:05")
	usr.Password = string(passHash)

	return usrUseCase.usrRepo.UpdateUser(usr)
}

func (usrUseCase *userUsecase) DeleteUser(username string) error {
	user , _:= usrUseCase.usrRepo.GetUserByUsername(username)
	if user == nil {
		return fmt.Errorf("user %v does not exist", username)
	}
	return usrUseCase.usrRepo.DeleteUser(username)
}

func NewUserUsecase(usrRepo repository.UserRepository) UserUsecase {
	return &userUsecase{
		usrRepo: usrRepo,
	}
}