package repository

import (
	"database/sql"
	"easylib-go/model"
	"easylib-go/utils"
	"fmt"
)

type MemberRepository interface {
	GetMemberById(string) (model.Member, error)
	InsertMember(model.Member) model.Member
	DeleteMember(model.Member) error
}

type memberRepository struct {
	db *sql.DB
}

func (mmbRepo *memberRepository) GetMemberById(id string) (model.Member, error) {
	qry := utils.GET_IMAGE_BY_ID
	mmb := model.Member{}
	err := mmbRepo.db.QueryRow(qry, id).Scan(&mmb.Id, &mmb.Name, &mmb.PhoneNo, &mmb.NoIdentity, &mmb.ImagePath, &mmb.LoanStatus, &mmb.CreatedAt, &mmb.UpdatedAt, &mmb.CreatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			panic(err)
		}
		panic(err)
	}
	return mmb, nil
}

func (mmbRepo *memberRepository) InsertMember(mmb model.Member) model.Member {
	qry := utils.INSERT_IMAGE
	_, err := mmbRepo.db.Exec(qry, &mmb.Id, &mmb.Name, &mmb.PhoneNo, &mmb.NoIdentity, &mmb.ImagePath, &mmb.LoanStatus, &mmb.CreatedAt, &mmb.UpdatedAt, &mmb.CreatedBy)
	if err != nil {
		panic(err)
	}
	return mmb
}

func (mmbRepo *memberRepository) DeleteMember(mmb model.Member) error {
	qry := utils.DELETE_IMAGE
	_, err := mmbRepo.db.Exec(qry, mmb.Id)
	if err != nil {
		return fmt.Errorf("error on memberRepository.DeleteMember : %v", err)
	}
	return nil
}

// func (mmbRepo *memberRepository) GetMemberByName(name string) (*model.Member, error) {
// 	qry := utils.GET_SERVICE_BY_NAME

// 	mmb := &model.Member{}
// 	err := mmbRepo.db.QueryRow(qry, name).Scan(&mmb.Id, &mmb.Name, &mmb.Uom, &mmb.Price)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, nil
// 		}
// 		return nil, fmt.Errorf("error on memberRepository.GetMemberByName() : %w", err)
// 	}
// 	return mmb, nil
// }

func NewMemberRepository(db *sql.DB) MemberRepository {
	return &memberRepository{
		db: db,
	}
}



