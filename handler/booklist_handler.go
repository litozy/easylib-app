package handler

import (
	"easylib-go/middleware"
	"easylib-go/model"
	"easylib-go/usecase"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	bkUsecase usecase.BookListUsecase
}

func (bkHandler *BookHandler) GetBookById(ctx *gin.Context) {
	idText := ctx.Param("id")
	if idText == "" {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"success":      false,
			"errorMessage": "Id must not be empty",
		})
		return
	}

	bk, err := bkHandler.bkUsecase.GetBookById(idText)
	if err != nil {
		fmt.Printf("bkHandler.GetBookById() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "There was an error getting the book data",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    bk,
	})
}

func (bkHandler *BookHandler) GetAllBook(ctx *gin.Context) {
	bk, err := bkHandler.bkUsecase.GetAllBook()
	if err != nil {
		fmt.Printf("bkHandler.bkUseCase.getAllBook() : %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "There was an error getting book data",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    bk,
	})
}

func (bkHandler *BookHandler) InsertBook(ctx *gin.Context) {
	bk := &model.Book{}
	err := ctx.ShouldBindJSON(&bk)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}
	err = bkHandler.bkUsecase.InsertBook(bk, ctx)
	if err != nil {
		fmt.Printf("bkHandler.InsertBook() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "There was error inserting book",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Book successfully inserted",
	})
}

func (bkHandler *BookHandler) DeleteBook(ctx *gin.Context) {
	idText := ctx.Param("id")
	if idText == "" {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"success":      false,
			"errorMessage": "Id must not be empty",
		})
		return
	}

	err := bkHandler.bkUsecase.DeleteBook(idText)
	if err != nil {
		fmt.Printf("bkHandler.bkUseCase.DeleteBook() : %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "There was error deleting book",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Book successfully deleted",
	})
}

func (bkHandler *BookHandler) UpdateBook(ctx *gin.Context) {
	bk := &model.Book{}
	bk.Id = ctx.Param("id")
	if bk.Id == "" {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"success":      false,
			"errorMessage": "Id must not be empty",
		})
		return
	}
	
	err := ctx.ShouldBindJSON(&bk)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}

	err = bkHandler.bkUsecase.UpdateBook(bk)
	if err != nil {
		fmt.Printf("bkHandler.bkUseCase.getAllBook() : %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan dalam memperbarui data Book",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Book successfully updated",
	})
}

func NewBookHandler(srv *gin.Engine, bkUsecase usecase.BookListUsecase) *BookHandler {
	bkHandler := &BookHandler{
		bkUsecase: bkUsecase,
	}
	srv.GET("/book/:id", middleware.RequireToken() ,bkHandler.GetBookById)
	srv.GET("/book", middleware.RequireToken(), bkHandler.GetAllBook)
	srv.POST("/book",middleware.RequireToken(), bkHandler.InsertBook)
	srv.DELETE("/book/:id", middleware.RequireToken(), bkHandler.DeleteBook)
	srv.PUT("/book/:id", middleware.RequireToken(), bkHandler.UpdateBook)
	return bkHandler
}
