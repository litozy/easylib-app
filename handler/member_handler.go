package handler

import (
	"easylib-go/middleware"
	"easylib-go/model"
	"easylib-go/usecase"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type MembersHandler struct {
	mmbUsecase usecase.MembersUsecase
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
		fmt.Printf("mmbHandler.InsertMember() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "There was error inserting member data",
		})
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
		fmt.Printf("mmbHandler.InsertMember() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "There was error inserting member data",
		})
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
		fmt.Printf("imgHandler.GetImageById() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "There was an error getting the image data",
		})
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
		fmt.Printf("mmbHandler.GetMemberById() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "There was an error getting the member data",
		})
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

    mmbHandler.mmbUsecase.DeleteMember(idText)
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Member successfully deleted",
	})
}

func NewMemberHandler(srv *gin.Engine, mmbUsecase usecase.MembersUsecase) *MembersHandler {
	mmbHandler := &MembersHandler{mmbUsecase: mmbUsecase}
	srv.GET("/memberimg/:id", middleware.RequireToken(), mmbHandler.GetMemberImageById)
	srv.POST("/member", middleware.RequireToken(), mmbHandler.InsertMember)
	srv.DELETE("/member/:id", middleware.RequireToken(), mmbHandler.DeleteMember)
	srv.GET("/member/:id", middleware.RequireToken(), mmbHandler.GetMemberDataById)
	srv.PUT("/member", middleware.RequireToken(), mmbHandler.UpdateMember)
	return mmbHandler

}