package model

import "time"

type Member struct {
	Id         string
	PhoneNo    string
	NoIdentity string
	Photo      []byte
	LoanStatus string
	CreatedAt  time.Time
	CreatedBy  string
}