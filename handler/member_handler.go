package handler

import (
	"easylib-go/model"
	"easylib-go/usecase"
	"easylib-go/utils"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type MembersHandler struct {
	mmbUsecase usecase.MembersUsecase
}

func (mmbHandler *MembersHandler) GetAllMembers(ctx *gin.Context) {
	mmb, err := mmbHandler.mmbUsecase.GetAllMember()
	if err != nil {
		fmt.Printf("mmbHandler.mmbUseCase.getAllMember() : %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "There was an error while getting the member data",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    mmb,
	})
}

func (mmbHandler *MembersHandler) InsertMember(ctx *gin.Context) {
	ctx.Request.ParseMultipartForm(10 * 1024 * 1024)

	req := model.MemberCreateRequest{} 
	req.FormData = ctx.Request.MultipartForm.File["image"]

	mmb := &model.Member{}
	err := ctx.ShouldBind(&mmb)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}
	err = mmbHandler.mmbUsecase.InsertMember(mmb, ctx, req)
	if err != nil {
		appError := &utils.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("mmbHandler.mmbUsecase.InsertMember() 1: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("mmbHandler.mmbUsecase.InsertMember() 2: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "An error occurred while inserting member data",
			})
		}
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Member successfully inserted",
	})
}

func (mmbHandler *MembersHandler) UpdateMember(ctx *gin.Context) {
	ctx.Request.ParseMultipartForm(10 * 1024 * 1024)

	req := model.MemberCreateRequest{}
	req.FormData = ctx.Request.MultipartForm.File["image"]

	mmb := &model.Member{}
	err := ctx.ShouldBind(&mmb)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}
	err = mmbHandler.mmbUsecase.UpdateMember(mmb, ctx, req)
	if err != nil {
		appError := &utils.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("mmbHandler.mmbUsecase.UpdateMember() 1: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("mmbHandler.mmbUsecase.UpdateMember() 2: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "An error occurred while updating member data",
			})
		}
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Member successfully inserted",
	})
}

func (mmbHandler *MembersHandler) GetMemberImageById(ctx *gin.Context) {
	idText := ctx.Param("id")
	if idText == "" {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"success":      false,
			"errorMessage": "Id must not be empty",
		})
		return
	}

	exampleResponse, _ := mmbHandler.mmbUsecase.GetMemberById(idText)
	fileBytes, err := os.ReadFile("public/" + exampleResponse.ImagePath)
	if err != nil {
		appError := &utils.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("mmbHandler.mmbUsecase.GetMemberImageById() 1: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("mmbHandler.mmbUsecase.GetMemberImageById() 2: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "An error occurred while getting member image",
			})
		}
		return
	}
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"success": true,
	// 	"data": exampleResponse,
	// })
	ctx.Writer.Write(fileBytes)
} 

func (mmbHandler *MembersHandler) GetMemberDataById(ctx *gin.Context) {
	idText := ctx.Param("id")
	if idText == "" {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"success":      false,
			"errorMessage": "Id must not be empty",
		})
		return
	}

	mmb, err := mmbHandler.mmbUsecase.GetMemberById(idText)
	if err != nil {
		appError := &utils.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("mmbHandler.mmbUsecase.GetMemberById() 1: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("mmbHandler.mmbUsecase.GetMemberMemberById() 2: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "An error occurred while getting member data",
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    mmb,
	})

} 

func (mmbHandler *MembersHandler) DeleteMember(ctx *gin.Context) {
	idText := ctx.Param("id")
	if idText == "" {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"success":      false,
			"errorMessage": "Id must not be empty",
		})
		return
	}

    err := mmbHandler.mmbUsecase.DeleteMember(idText)
	if err != nil {
		appError := &utils.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("mmbHandler.mmbUsecase.DeleteMember() 1: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("mmbHandler.mmbUsecase.DeleteMember() 2: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "An error occurred while deleting member data",
			})
		}
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Member successfully deleted",
	})
}

func NewMemberHandler(srv *gin.Engine, mmbUsecase usecase.MembersUsecase) *MembersHandler {
	mmbHandler := &MembersHandler{mmbUsecase: mmbUsecase}
	srv.GET("/member", mmbHandler.GetAllMembers)
	srv.GET("/memberimg/:id", mmbHandler.GetMemberImageById)
	srv.POST("/member", mmbHandler.InsertMember)
	srv.DELETE("/member/:id", mmbHandler.DeleteMember)
	srv.GET("/member/:id", mmbHandler.GetMemberDataById)
	srv.PUT("/member", mmbHandler.UpdateMember)
	return mmbHandler

}

// middleware.RequireToken()