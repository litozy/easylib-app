package model

import "time"

type BookLoan struct {
	Id        string
	MemberId  string
	BookId    []string
	StartDate time.Time
	EndDate time.Time
	LateChargeDay int
	LateCharge float64
	LoanStatus string
}

type BookLoanView struct {
	Name string
	LoanStatus string
	LateChargeTotal float64
	Loaning []BookLoanViewDetail
}

type BookLoanViewDetail struct {
	BookName string
	StartDate time.Time
	EndDate time.Time
	LateChargeDay int
	LateCharge float64
	LoanStatus string
}