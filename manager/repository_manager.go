package manager

import (
	"easylib-go/repository"
	"sync"
)

type RepositoryManager interface {
	GetUserRepository() repository.UserRepository
}

type repositoryManager struct {
	infraManager InfraManager
	usrRepo repository.UserRepository
}

var onceLoadUserRepo sync.Once

func (rm *repositoryManager) GetUserRepository() repository.UserRepository {
	onceLoadUserRepo.Do(func() {
		rm.usrRepo = repository.NewUserRepository(rm.infraManager.GetDB())
	})
	return rm.usrRepo
}

func NewRepoManager(infraManager InfraManager) RepositoryManager {
	return &repositoryManager{
		infraManager: infraManager,
	}
}