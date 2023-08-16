package model

import (
	"mime/multipart"
)

type Images struct {
	Id        string
	Path      string
	CreatedAt string
	UpdatedAt string
}

type ImageCreateRequest struct {
	FormData []*multipart.FileHeader
}

type ImageResponse struct {
	Id        string `json:"id"`
	Path      string `json:"path"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}