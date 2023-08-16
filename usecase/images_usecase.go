package usecase

import (
	"easylib-go/model"
	"easylib-go/repository"
	"easylib-go/utils"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type ImagesUsecase interface {
	InsertImage(model.ImageCreateRequest) []model.ImageResponse
	DeleteImage(string) error
	GetImageById(string) (*model.Images, error)
}

type imagesUsecase struct {
	imgRepo repository.ImagesRepository
}

func (imgUsecase *imagesUsecase) InsertImage(req model.ImageCreateRequest) []model.ImageResponse {
	var imageResponses []model.ImageResponse

	for _, image := range req.FormData {
		file, _ := image.Open()

		tempFile, err := os.CreateTemp("public", "image-*.jpg")
		if err != nil {
			panic(err)
		}
		defer tempFile.Close()

		fileBytes, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}

		tempFile.Write(fileBytes)

		fileName := tempFile.Name()
		newFileName := strings.Split(fileName, "\\")

		CreatedAt := time.Now().UTC()
		UpdatedAt := time.Now().UTC()
		image := model.Images{
			Id:        utils.UuidGenerate(),
			Path:      newFileName[1],
			CreatedAt: CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: UpdatedAt.Format("2006-01-02 15:04:05"),
		}

		image = imgUsecase.imgRepo.InsertImage(image)
		imageResponses = append(imageResponses, utils.ToImageResponse(image))
	}
	return imageResponses
} 

func (imgUsecase *imagesUsecase) DeleteImage(id string) error {
	image, _ := imgUsecase.imgRepo.GetImageById(id)
	if image == nil {
		return fmt.Errorf("user %v does not exist", id)
	}
	os.Remove("public/" + image.Path)

	return imgUsecase.imgRepo.DeleteImage(id)
}

func (imgUsecase *imagesUsecase) GetImageById(id string) (*model.Images, error) {
	return imgUsecase.imgRepo.GetImageById(id)
}

func NewImagesUsecase(imgRepo repository.ImagesRepository) ImagesUsecase {
	return &imagesUsecase{imgRepo: imgRepo}
}