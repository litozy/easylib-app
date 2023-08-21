package model

type BookLoan struct {
	Id         string
	MemberId   string
	BookId     []string
	StartDate  string
	EndDate    string
	LoanStatus string
}

type BookLoanView struct {
	Id         string
	Name       string
	LoanStatus string
	Loaning    []BookLoanViewDetail
}

type BookLoanViewDetail struct {
	BookName   string
	StartDate  string
	EndDate    string
	LoanStatus string
}