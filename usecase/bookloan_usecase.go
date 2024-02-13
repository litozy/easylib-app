package usecase

import (
	"easylib-go/model"
	"easylib-go/repository"
	"easylib-go/utils"
	"fmt"
	"time"
)

type BookLoanUsecase interface {
	InsertBookLoan(bl *model.BookLoan) error
	GetBookLoanByMemberId(memberId string) (*model.BookLoanView, error)
	GetAllBookLoan() ([]model.BookLoanView, error)
	UpdateBookLoan(bl *model.BookLoan) error
}

type bookLoanUsecase struct {
	mmbRepo repository.MemberRepository
	blRepo repository.BookLoan
}

func (blUsecase *bookLoanUsecase) InsertBookLoan(bl *model.BookLoan) error {
	member, _ := blUsecase.mmbRepo.GetMemberById(bl.MemberId)
	if member == nil {
		return &utils.AppError{
			ErrorCode: 1,
			ErrorMessage: fmt.Sprintf("there are no member with id : %s", bl.MemberId),
		}
	}
	startDate := time.Now().UTC()
	bl.StartDate = startDate.Format("2006-01-02")
	bl.LoanStatus = true
	
	return blUsecase.blRepo.InsertBookLoan(bl)
}

func (blUsecase *bookLoanUsecase) GetBookLoanByMemberId(memberId string) (*model.BookLoanView, error) {
	return blUsecase.blRepo.GetBookLoanByMemberId(memberId)
}

func (blUsecase *bookLoanUsecase) GetAllBookLoan() ([]model.BookLoanView, error) {
	return blUsecase.blRepo.GetAllLoanBooks()
}

func (blUsecase *bookLoanUsecase) UpdateBookLoan(bl *model.BookLoan) error {
	return blUsecase.blRepo.UpdateBookLoan(bl)
}

func NewBookLoanUsecase(blRepo repository.BookLoan, mmbRepo repository.MemberRepository) BookLoanUsecase {
	return &bookLoanUsecase{blRepo: blRepo, mmbRepo: mmbRepo}
}