package manager

import (
	"easylib-go/repository"
	"sync"
)

type RepositoryManager interface {
	GetUserRepository() repository.UserRepository
	GetBookRepository() repository.BookListRepository
	GetMemberRepository() repository.MemberRepository
	GetBookLoanRepository() repository.BookLoan
}

type repositoryManager struct {
	infraManager InfraManager
	usrRepo repository.UserRepository
	bkRepo repository.BookListRepository
	mmbRepo repository.MemberRepository
	blRepo repository.BookLoan
}

var onceLoadUserRepo sync.Once
var onceLoadBookRepo sync.Once
var onceLoadMemberRepo sync.Once
var onceLoadBookLoanRepo sync.Once

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

func (rm *repositoryManager) GetBookLoanRepository() repository.BookLoan {
	onceLoadBookLoanRepo.Do(func() {
		rm.blRepo = repository.NewBookLoanRepository(rm.infraManager.GetDB())
	})
	return rm.blRepo
}


func NewRepoManager(infraManager InfraManager) RepositoryManager {
	return &repositoryManager{
		infraManager: infraManager,
	}
}