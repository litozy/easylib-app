package model

import (
	"mime/multipart"
)

type Member struct {
	Id         string
	Name 	   string
	PhoneNo    string
	NoIdentity string
	ImagePath  string
	LoanStatus string
	CreatedAt  string
	UpdatedAt  string
	CreatedBy  string
}

type MemberCreateRequest struct {
	FormData []*multipart.FileHeader
}

type MemberResponse struct {
	Id         string `json:"id"`
	Name 	   string `json:"name"`
	PhoneNo    string `json:"phone_no"`
	NoIdentity string `json:"no_identity"`
	ImagePath  string `json:"image_path"`
	LoanStatus string `json:"loan_status"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	CreatedBy  string `json:"created_by"`
}