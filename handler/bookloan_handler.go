package handler

import (
	"easylib-go/middleware"
	"easylib-go/model"
	"easylib-go/usecase"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookLoanHandler struct {
	blUsecase usecase.BookLoanUsecase
}

func (blHandler *BookLoanHandler) InsertBookLoan(ctx *gin.Context) {
	bl := &model.BookLoan{}
	err := ctx.ShouldBindJSON(&bl)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}
	err = blHandler.blUsecase.InsertBookLoan(bl)
	if err != nil {
		fmt.Printf("blHandler.InsertBookLoan() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "There was error inserting BookLoan",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Book loaning successfully inserted",
	})
}

func (blHandler *BookLoanHandler) GetBookLoanById(ctx *gin.Context) {
	idText := ctx.Param("id")
	if idText == "" {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"success":      false,
			"errorMessage": "Id must not be empty",
		})
		return
	}

	bl, err := blHandler.blUsecase.GetBookLoanByMemberId(idText)
	if err != nil {
		fmt.Printf("blHandler.GetBookLoanById() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "There was an error getting the BookLoan data",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    bl,
	})
}

func (blHandler *BookLoanHandler) GetAllBookLoan(ctx *gin.Context) {
	bl, err := blHandler.blUsecase.GetAllBookLoan()
	if err != nil {
		fmt.Printf("blHandler.blUseCase.getAllBookLoan() : %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "There was an error getting BookLoan data",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    bl,
	})
}

func (blHandler *BookLoanHandler) UpdateBookLoan(ctx *gin.Context) {
	bl := &model.BookLoan{}
	// bl.Id = ctx.Param("id")
	// if bl.Id == "" {
	// 	ctx.JSON(http.StatusBadGateway, gin.H{
	// 		"success":      false,
	// 		"errorMessage": "Id must not be empty",
	// 	})
	// 	return
	// }

	err := ctx.ShouldBindJSON(&bl)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}

	err = blHandler.blUsecase.UpdateBookLoan(bl)
	if err != nil {
		fmt.Printf("blHandler.blUseCase.UpdateBookLoan() : %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan dalam memperbarui data BookLoan",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "BookLoan successfully updated",
	})
}

func NewBookLoanHandler(srv *gin.Engine, blUsecase usecase.BookLoanUsecase) *BookLoanHandler {
	blHandler := &BookLoanHandler{
		blUsecase: blUsecase,
	}
	srv.GET("/bookloan/:id", middleware.RequireToken() ,blHandler.GetBookLoanById)
	srv.GET("/bookloan", middleware.RequireToken(), blHandler.GetAllBookLoan)
	srv.POST("/bookloan",middleware.RequireToken(), blHandler.InsertBookLoan)
	srv.PUT("/bookloan", middleware.RequireToken(), blHandler.UpdateBookLoan)

	return blHandler
}