package handler

import (
	"easylib-go/model"
	"easylib-go/usecase"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type MembersHandler struct {
	imgUsecase usecase.MembersUsecase
}

func (imgHandler *MembersHandler) InsertMember(ctx *gin.Context) {
	mem := &model.Member{}
	err := ctx.ShouldBindJSON(&mem)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}
	ctx.Request.ParseMultipartForm(10 * 1024 * 1024)

	memberCreateRequest := model.MemberCreateRequest{}
	memberCreateRequest.FormData = ctx.Request.MultipartForm.File["member"]

	imgHandler.imgUsecase.InsertMember(ctx, mem, memberCreateRequest)
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Member inserted successfully",
	})
}

func (imgHandler *MembersHandler) DeleteMember(ctx *gin.Context) {
	idText := ctx.Param("id")
	if idText == "" {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"success":      false,
			"errorMessage": "Id must not be empty",
		})
		return
	}

    imgHandler.imgUsecase.DeleteMember(idText)
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Member successfully deleted",
	})
}

func (imgHandler *MembersHandler) GetMemberById(ctx *gin.Context) {
	idText := ctx.Param("id")
	if idText == "" {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"success":      false,
			"errorMessage": "Id must not be empty",
		})
		return
	}

	exampleResponse := imgHandler.imgUsecase.GetMemberById(idText)
	fileBytes, err := os.ReadFile("public/" + exampleResponse.ImagePath)
	if err != nil {
		fmt.Printf("imgHandler.GetMemberById() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "There was an error getting the member data",
		})
		return 
	}
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"success": true,
	// 	"data": exampleResponse,
	// })
	ctx.Writer.Write(fileBytes)
} 

func NewMemberHandler(srv *gin.Engine, imgUsecase usecase.MembersUsecase) *MembersHandler {
	imgHandler := &MembersHandler{imgUsecase: imgUsecase}
	srv.GET("/member/:id", imgHandler.GetMemberById)
	srv.POST("/member", imgHandler.InsertMember)
	srv.DELETE("/member/:id", imgHandler.DeleteMember)
	return imgHandler

}