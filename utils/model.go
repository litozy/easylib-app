package utils

import "easylib-go/model"

func ToImageResponse(image model.Images) model.ImageResponse {
	return model.ImageResponse{
		Id:        image.Id,
		Path:      image.Path,
		CreatedAt: image.CreatedAt,
		UpdatedAt: image.UpdatedAt,
	}
}