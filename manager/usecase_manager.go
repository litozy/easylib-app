package manager

import (
	"easylib-go/usecase"
	"sync"
)

type UsecaseManager interface {
	GetUserUsecase() usecase.UserUsecase
	GetLoginUsecase() usecase.LoginUseCase
	GetBookListUsecase() usecase.BookListUsecase
	GetMemberUsecase() usecase.MembersUsecase
}

type usecaseManager struct {
	repositoryManager RepositoryManager
	usrUsecase        usecase.UserUsecase
	lgUsecase     usecase.LoginUseCase
	bkUsecase  	usecase.BookListUsecase
	mmbUsecase usecase.MembersUsecase
}

var onceLoadUserUsecase sync.Once
var onceLoadLoginUsecase sync.Once
var onceLoadBookListUsecase sync.Once
var onceLoadImageUsecase sync.Once
var onceLoadMemberUsecase sync.Once


func (um *usecaseManager) GetUserUsecase() usecase.UserUsecase {
	onceLoadUserUsecase.Do(func() {
		um.usrUsecase = usecase.NewUserUsecase(um.repositoryManager.GetUserRepository())
	})
	return um.usrUsecase
}

func (um *usecaseManager) GetLoginUsecase() usecase.LoginUseCase {
	onceLoadLoginUsecase.Do(func() {
		um.lgUsecase = usecase.NewLoginUseCase(um.repositoryManager.GetUserRepository())
	})
	return um.lgUsecase
}

func (um *usecaseManager) GetBookListUsecase() usecase.BookListUsecase {
	onceLoadBookListUsecase.Do(func() {
		um.bkUsecase = usecase.NewBookUsecase(um.repositoryManager.GetBookRepository())

	})
	return um.bkUsecase
}

func (um *usecaseManager) GetMemberUsecase() usecase.MembersUsecase {
	onceLoadMemberUsecase.Do(func() {
		um.mmbUsecase = usecase.NewMembersUsecase(um.repositoryManager.GetMemberRepository())

	})
	return um.mmbUsecase
}

func NewUsecaseManager(repositoryManager RepositoryManager) UsecaseManager {
	return &usecaseManager{
		repositoryManager: repositoryManager,
	}
}