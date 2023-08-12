package model

import "time"

type Book struct {
	Id        string
	BookName  string
	CreatedAt time.Time
	CreatedBy string
	Stock     int
}