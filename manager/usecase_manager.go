package manager

import (
	"easylib-go/usecase"
	"sync"
)

type UsecaseManager interface {
	GetUserUsecase() usecase.UserUsecase
}

type usecaseManager struct {
	repositoryManager RepositoryManager
	usrUsecase        usecase.UserUsecase
}

var onceLoadUserUsecase sync.Once

func (um *usecaseManager) GetUserUsecase() usecase.UserUsecase {
	onceLoadUserUsecase.Do(func() {
		um.usrUsecase = usecase.NewUserUsecase(um.repositoryManager.GetUserRepository())
	})
	return um.usrUsecase
}

func NewUsecaseManager(repositoryManager RepositoryManager) UsecaseManager {
	return &usecaseManager{
		repositoryManager: repositoryManager,
	}
}