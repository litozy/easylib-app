package manager

import (
	"easylib-go/repository"
	"sync"
)

type RepositoryManager interface {
	GetUserRepository() repository.UserRepository
	GetBookRepository() repository.BookListRepository
}

type repositoryManager struct {
	infraManager InfraManager
	usrRepo repository.UserRepository
	bkRepo repository.BookListRepository
}

var onceLoadUserRepo sync.Once
var onceLoadBookRepo sync.Once

func (rm *repositoryManager) GetUserRepository() repository.UserRepository {
	onceLoadUserRepo.Do(func() {
		rm.usrRepo = repository.NewUserRepository(rm.infraManager.GetDB())
	})
	return rm.usrRepo
}

func (rm *repositoryManager) GetBookRepository() repository.BookListRepository {
	onceLoadBookRepo.Do(func() {
		rm.bkRepo = repository.NewBookRepo(rm.infraManager.GetDB())
	})
	return rm.bkRepo
}

func NewRepoManager(infraManager InfraManager) RepositoryManager {
	return &repositoryManager{
		infraManager: infraManager,
	}
}