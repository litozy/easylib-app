package repository

import (
	"database/sql"
	"easylib-go/model"
	"easylib-go/utils"
	"fmt"
)

type BookLoan interface {
	InsertBookLoan(bl *model.BookLoan) error
	UpdateBookLoan(bl *model.BookLoan) error
	GetBookLoanByMemberId(memberId string) (*model.BookLoanView, error)
	GetAllLoanBooks() ([]model.BookLoanView, error)
}

type bookLoan struct {
	db *sql.DB
}

func (blRepo *bookLoan) InsertBookLoan(bl *model.BookLoan) error {
	bk := &model.Book{}
	tx, err := blRepo.db.Begin()
	if err != nil {
		return fmt.Errorf("blRepo.InsertBookLoan() Begin: %w", err)
	}

	qry := utils.INSERT_BOOKLOAN
	_, err = tx.Exec(qry, bl.Id, bl.MemberId, bl.BookId, bl.StartDate, bl.EndDate, bl.LoanStatus)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("blRepo.InsertBookLoan() Exec: %w", err)
	}

	// Update member's loan status
	qry = utils.UPDATE_MEMBER_BOOKLOAN
	_, err = tx.Exec(qry, true, bl.MemberId)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("blRepo.InsertBookLoan() Update Member Loan Status: %w", err)
	}

	// Decrease book stock
	qry = utils.UPDATE_STOCK_BOOKLOAN
	_, err = tx.Exec(qry, bk.UpdatedAt, bl.BookId)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("blRepo.InsertBookLoan() Update Book Stock: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("blRepo.InsertBookLoan() Commit: %w", err)
	}

	return nil
}

func (blRepo *bookLoan) GetBookLoanByMemberId(memberId string) (*model.BookLoanView, error) {
	qry := utils.GET_BOOKLOAN_BY_ID
	
	bl := &model.BookLoanView{}
	err := blRepo.db.QueryRow(qry, memberId).Scan(&bl.Id, &bl.Name, &bl.LoanStatus)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("error on blRepo.GetBookLoanByMemberId() 1: %v", err)
		}
		return nil, err
	}
	
	qryDet := utils.GET_BOOKLOAN_BY_ID_DETAIL
	
	rows, err := blRepo.db.Query(qryDet, memberId)
	if err != nil {
		return nil, fmt.Errorf("error on blRepo.GetBookLoanByMemberId() 2: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var bookName, startDate, endDate, loanStatus string
		if err := rows.Scan(&bookName, &startDate, &endDate, &loanStatus); err != nil {
			return nil, fmt.Errorf("error on blRepo.GetBookLoanByMemberId() 3: %v", err)
		}

		loaningItem := model.BookLoanViewDetail{
			BookName:   bookName,
			StartDate:  startDate,
			EndDate:    endDate,
			LoanStatus: loanStatus,
		}
		bl.Loaning = append(bl.Loaning, loaningItem)
	}

	return bl, nil
}

func (blRepo *bookLoan) GetAllLoanBooks() ([]model.BookLoanView, error) {
	qry := utils.GET_ALL_BOOKLOAN
	
	rows, err := blRepo.db.Query(qry)
	if err != nil {
		return nil, fmt.Errorf("error on blRepo.GetAllLoanBooks() 1: %v", err)
	}
	defer rows.Close()

	loanBooks := make([]model.BookLoanView, 0)
	currentMemberID := ""
	loanBook := model.BookLoanView{}

	for rows.Next() {
		var memberID, memberName, memberLoanStatus, bookName, startDate, endDate, loanStatus string
		if err := rows.Scan(&memberID, &memberName, &memberLoanStatus, &bookName, &startDate, &endDate, &loanStatus); err != nil {
			return nil, fmt.Errorf("error on blRepo.GetAllLoanBooks() 2: %v", err)
		}

		if memberID != currentMemberID {
			if currentMemberID != "" {
				loanBooks = append(loanBooks, loanBook)
			}

			loanBook = model.BookLoanView{
				Id:         memberID,
				Name:       memberName,
				LoanStatus: memberLoanStatus,
				Loaning:    make([]model.BookLoanViewDetail, 0),
			}

			currentMemberID = memberID
		}

		loaningItem := model.BookLoanViewDetail{
			BookName:   bookName,
			StartDate:  startDate,
			EndDate:    endDate,
			LoanStatus: loanStatus,
		}
		loanBook.Loaning = append(loanBook.Loaning, loaningItem)
	}

	if currentMemberID != "" {
		loanBooks = append(loanBooks, loanBook)
	}

	return loanBooks, nil
}

func (blRepo *bookLoan) UpdateBookLoan(bl *model.BookLoan) error {
	qry := utils.UPDATE_BOOKLOAN_STATUS
	_, err := blRepo.db.Exec(qry, &bl.Id, &bl.LoanStatus)
	if err != nil {
		return fmt.Errorf("error on blRepository.UpdateBookLoan : %v", &err)
	}
	return nil
}

func NewBookLoanRepository(db *sql.DB) BookLoan {
	return &bookLoan{db: db}
}