package repository

import (
	"database/sql"
	"easylib-go/model"
	"easylib-go/utils"
	"fmt"
)

type BookListRepository interface {
	GetBookById(string) (*model.Book, error)
	GetAllBook() ([]*model.Book, error)
	InsertBook(*model.Book) error
	DeleteBook(string) error
	UpdateBook(*model.Book) error
}

type bookListRepository struct {
	db *sql.DB
}

func (bkRepo *bookListRepository) GetBookById(id string) (*model.Book, error) {
	qry := utils.GET_BOOK_BY_ID
	bk := &model.Book{}
	err := bkRepo.db.QueryRow(qry, id).Scan(&bk.Id, &bk.BookName, &bk.CreatedAt, &bk.CreatedBy, &bk.Stock)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on bookListRepository.getBookById() : %v", err)
	}
	return bk, nil
}

func (bkRepo *bookListRepository) GetAllBook() ([]*model.Book, error) {
	qry := utils.GET_ALL_BOOK
	var arrBook []*model.Book
	rows, err := bkRepo.db.Query(qry)
	if err != nil {
		return nil, fmt.Errorf("getAllBook error : %v", err)
	}

	for rows.Next() {
		bk := &model.Book{}
		rows.Scan(&bk.Id, &bk.BookName, &bk.CreatedAt, &bk.CreatedBy, &bk.Stock)
		arrBook = append(arrBook, bk)
	}
	return arrBook, nil

}

func (bkRepo *bookListRepository) InsertBook(bk *model.Book) error {
	qry := utils.INSERT_BOOK
	_, err := bkRepo.db.Exec(qry, bk.Id, bk.BookName, bk.CreatedAt, bk.CreatedBy, bk.Stock)
	if err != nil {
		return fmt.Errorf("error on bookListRepository.InsertBook() : %w", err)
	}
	return nil
}

func (bkRepo *bookListRepository) DeleteBook(id string) error {
	qry := utils.DELETE_BOOK
	_, err := bkRepo.db.Exec(qry, id)
	if err != nil {
		return fmt.Errorf("error on bookListRepository.DeleteBook : %v", err)
	}
	return nil
}

func (bkRepo *bookListRepository) UpdateBook(bk *model.Book) error {
	qry := utils.UPDATE_BOOK_STOCK
	_, err := bkRepo.db.Exec(qry, bk.Stock)
	if err != nil {
		return fmt.Errorf("error on bookListRepository.UpdateBook : %v", &err)
	}
	return nil
}

// func (bkRepo *bookListRepository) GetBookByName(name string) (*model.Book, error) {
// 	qry := utils.GET_SERVICE_BY_NAME

// 	bk := &model.Book{}
// 	err := bkRepo.db.QueryRow(qry, name).Scan(&bk.Id, &bk.Name, &bk.Uom, &bk.Price)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, nil
// 		}
// 		return nil, fmt.Errorf("error on bookListRepository.GetBookByName() : %w", err)
// 	}
// 	return bk, nil
// }

func NewBookRepository(db *sql.DB) BookListRepository {
	return &bookListRepository{
		db: db,
	}
}


