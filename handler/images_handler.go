package handler

import (
	"easylib-go/model"
	"easylib-go/usecase"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type ImagesHandler struct {
	imgUsecase usecase.ImagesUsecase
}

func (imgHandler *ImagesHandler) InsertImage(ctx *gin.Context) {
	ctx.Request.ParseMultipartForm(10 * 1024 * 1024)

	imageCreateRequest := model.ImageCreateRequest{}
	imageCreateRequest.FormData = ctx.Request.MultipartForm.File["image"]

	imgHandler.imgUsecase.InsertImage(imageCreateRequest)
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Image inserted successfully",
	})
}

func (imgHandler *ImagesHandler) DeleteImage(ctx *gin.Context) {
	idText := ctx.Param("id")
	if idText == "" {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"success":      false,
			"errorMessage": "Id must not be empty",
		})
		return
	}

	err := imgHandler.imgUsecase.DeleteImage(idText)
	if err != nil {
		fmt.Printf("imgHandler.imgUseCase.DeleteImage() : %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "There was error deleting image",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Image successfully deleted",
	})
}

func (imgHandler *ImagesHandler) GetImageById(ctx *gin.Context) {
	idText := ctx.Param("id")
	if idText == "" {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"success":      false,
			"errorMessage": "Id must not be empty",
		})
		return
	}

	exampleResponse, _ := imgHandler.imgUsecase.GetImageById(idText)
	fileBytes, err := os.ReadFile("public/" + exampleResponse.Path)
	if err != nil {
		fmt.Printf("imgHandler.GetImageById() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "There was an error getting the image data",
		})
		return 
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": fileBytes,
	})
} 

func NewImageHandler(srv *gin.Engine, imgUsecase usecase.ImagesUsecase) *ImagesHandler {
	imgHandler := &ImagesHandler{imgUsecase: imgUsecase}
	srv.GET("/image", imgHandler.GetImageById)
	srv.POST("/image", imgHandler.InsertImage)
	srv.DELETE("/image/:id", imgHandler.DeleteImage)
	return imgHandler

}