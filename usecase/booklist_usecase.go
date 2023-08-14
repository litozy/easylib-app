package usecase

import (
	"easylib-go/model"
	"easylib-go/repository"
	"easylib-go/utils"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type BookListUsecase interface {
	GetBookById(string) (*model.Book, error)
	GetAllBook() ([]*model.Book, error)
	InsertBook(*model.Book, *gin.Context) error
	DeleteBook(string) error
	UpdateBook(*model.Book) error
}
type bookListUsecase struct {
	bkRepo repository.BookListRepository
}

func (bkUsecase *bookListUsecase) GetBookById(id string) (*model.Book, error) {
	return bkUsecase.bkRepo.GetBookById(id)
}

func (bkUsecase *bookListUsecase) GetAllBook() ([]*model.Book, error) {
	return bkUsecase.bkRepo.GetAllBook()
}

func (bkUsecase *bookListUsecase) InsertBook(bk *model.Book, ctx *gin.Context) error {
	session := sessions.Default(ctx)
	existSession := session.Get("Username")

	bk.Id = utils.UuidGenerate()
	bk.CreatedAt = time.Now().UTC()
	bk.CreatedBy = existSession.(string)
	return bkUsecase.bkRepo.InsertBook(bk)
}

func (bkUsecase *bookListUsecase) DeleteBook(id string) error {
	return bkUsecase.bkRepo.DeleteBook(id)
}

func (bkUsecase *bookListUsecase) UpdateBook(bk *model.Book) error {
	return bkUsecase.bkRepo.UpdateBook(bk)
}

func NewBookUsecase(bkRepo repository.BookListRepository) BookListUsecase {
	return &bookListUsecase{
		bkRepo: bkRepo,
	}
}
