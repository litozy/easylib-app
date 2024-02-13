package repository

import (
	"database/sql"
	"easylib-go/model"
	"easylib-go/utils"
	"fmt"
	"time"
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

	stmt, err := tx.Prepare(utils.INSERT_BOOKLOAN)
    if err != nil {
        return err
    }
    defer stmt.Close()
	
	
	for _, bookId := range bl.BookId {
		bl.Id = utils.UuidGenerate()
		fmt.Println(bookId)
        _, err := stmt.Exec(bl.Id, bl.MemberId, bookId, bl.StartDate, bl.EndDate, bl.LoanStatus)
        if err != nil {
			tx.Rollback()
            return err
        }

		updatedAt := time.Now().UTC()
		bk.UpdatedAt = updatedAt.Format("2006-01-02")
		qry := utils.UPDATE_STOCK_BOOKLOAN
		_, err = tx.Exec(qry, bookId, bk.UpdatedAt)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("blRepo.InsertBookLoan() Update Book Stock: %w", err)
		}
    }

	// qry := utils.INSERT_BOOKLOAN
	// _, err = tx.Exec(qry, bl.Id, bl.MemberId, pq.Array(bl.BookId), bl.StartDate, bl.EndDate, bl.LoanStatus)
	// if err != nil {
	// 	tx.Rollback()
	// 	return fmt.Errorf("blRepo.InsertBookLoan() Exec: %w", err)
	// }

	// Update member's loan status
	qry := utils.UPDATE_MEMBER_BOOKLOAN
	_, err = tx.Exec(qry, bl.MemberId, true)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("blRepo.InsertBookLoan() Update Member Loan Status: %w", err)
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
		var memberID, memberName, memberLoanStatus, bookLoanId, bookName, startDate, endDate, loanStatus string
		if err := rows.Scan(&memberID, &memberName, &memberLoanStatus, &bookLoanId, &bookName, &startDate, &endDate, &loanStatus); err != nil {
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
			Id: 		bookLoanId,
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
    _, err := blRepo.db.Exec(qry, bl.Id, false)
    if err != nil {
        return fmt.Errorf("error on blRepository.UpdateBookLoan : %v", err)
    }

    qry2 := utils.GET_BOOKLOAN_STATUS_BY_ID
    rows, err := blRepo.db.Query(qry2, bl.Id)
    if err != nil {
        return fmt.Errorf("error on blRepository.UpdateMemberBookLoan: %v", err)
    }
    defer rows.Close()

	

    allLoansDone := false // Tentukan semua pinjaman buku selesai secara default
    for rows.Next() {
		
        if err := rows.Scan(&bl.Id, &bl.MemberId, &bl.LoanStatus); err != nil {
            return fmt.Errorf("error scanning row: %v", err)
        }

		qry3 := utils.GET_BOOKLOAN_STATUS_BY_MEMBER_ID
		rows2, err := blRepo.db.Query(qry3, bl.MemberId)
    	if err != nil {
        	return fmt.Errorf("error on blRepository.UpdateMemberBookLoan 2: %v", err)
    	}
    	defer rows.Close()

		for rows2.Next() {
			if err := rows2.Scan(&bl.Id, &bl.MemberId, &bl.LoanStatus); err != nil {
				return fmt.Errorf("error scanning row 2: %v", err)
			}

        	if bl.LoanStatus { // Jika ada pinjaman buku yang belum selesai, ubah status semua pinjaman buku dari memberID menjadi false
            	allLoansDone = true
            	break
        	}
		}
    }
	qry2 = utils.UPDATE_MEMBER_BOOKLOAN_DONE
    // Jika semua pinjaman buku dari memberID sudah selesai, ubah status pinjaman buku dari memberID menjadi false
    if !allLoansDone {
        _, err := blRepo.db.Exec(qry2, bl.MemberId, false)
        if err != nil {
            return fmt.Errorf("error on blRepository.UpdateMemberBookLoan Exec : %v", err)
        }
    }

    return nil
}


func NewBookLoanRepository(db *sql.DB) BookLoan {
	return &bookLoan{db: db}
}