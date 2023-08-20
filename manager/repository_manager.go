package manager

import (
	"easylib-go/repository"
	"sync"
)

type RepositoryManager interface {
	GetUserRepository() repository.UserRepository
	GetBookRepository() repository.BookListRepository
	GetMemberRepository() repository.MemberRepository
}

type repositoryManager struct {
	infraManager InfraManager
	usrRepo repository.UserRepository
	bkRepo repository.BookListRepository
	mmbRepo repository.MemberRepository
}

var onceLoadUserRepo sync.Once
var onceLoadBookRepo sync.Once
var onceLoadImageRepo sync.Once
var onceLoadMemberRepo sync.Once

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

func (rm *repositoryManager) GetMemberRepository() repository.MemberRepository {
	onceLoadMemberRepo.Do(func() {
		rm.mmbRepo = repository.NewMemberRepository(rm.infraManager.GetDB())
	})
	return rm.mmbRepo
}

func NewRepoManager(infraManager InfraManager) RepositoryManager {
	return &repositoryManager{
		infraManager: infraManager,
	}
}