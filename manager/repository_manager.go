package manager

import (
	"easylib-go/repository"
	"sync"
)

type RepositoryManager interface {
	GetUserRepository() repository.UserRepository
	GetBookRepository() repository.BookListRepository
	GetImageRepository() repository.ImagesRepository
}

type repositoryManager struct {
	infraManager InfraManager
	usrRepo repository.UserRepository
	bkRepo repository.BookListRepository
	imgRepo repository.ImagesRepository
}

var onceLoadUserRepo sync.Once
var onceLoadBookRepo sync.Once
var onceLoadImageRepo sync.Once

func (rm *repositoryManager) GetUserRepository() repository.UserRepository {
	onceLoadUserRepo.Do(func() {
		rm.usrRepo = repository.NewUserRepository(rm.infraManager.GetDB())
	})
	return rm.usrRepo
}

func (rm *repositoryManager) GetBookRepository() repository.BookListRepository {
	onceLoadBookRepo.Do(func() {
		rm.bkRepo = repository.NewBookRepository(rm.infraManager.GetDB())
	})
	return rm.bkRepo
}

func (rm *repositoryManager) GetImageRepository() repository.ImagesRepository {
	onceLoadImageRepo.Do(func() {
		rm.imgRepo = repository.NewImageRepository(rm.infraManager.GetDB())
	})
	return rm.imgRepo
}

func NewRepoManager(infraManager InfraManager) RepositoryManager {
	return &repositoryManager{
		infraManager: infraManager,
	}
}