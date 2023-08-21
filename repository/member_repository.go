package repository

import (
	"database/sql"
	"easylib-go/model"
	"easylib-go/utils"
	"fmt"
)

type MemberRepository interface {
	GetMemberById(string) (*model.Member, error)
	InsertMember(*model.Member) error
	DeleteMember(mmb *model.Member) error
	UpdateMember(*model.Member) error
	GetMemberByPhoneNumber(phoneno string) (*model.Member, error)
	GetMemberByIdMember(idmem string) (*model.Member, error)
	GetAllMember() ([]*model.Member, error)
}

type memberRepository struct {
	db *sql.DB
}

func (mmbRepo *memberRepository) GetMemberByPhoneNumber(phoneno string) (*model.Member, error) {
	qry := utils.GET_MEMBER_BY_PHONENO
	mmb := &model.Member{}
	err := mmbRepo.db.QueryRow(qry, phoneno).Scan(&mmb.Id, &mmb.Name, &mmb.PhoneNo, &mmb.NoIdentity, &mmb.ImagePath, &mmb.LoanStatus, &mmb.CreatedAt, &mmb.UpdatedAt, &mmb.CreatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("error on memberRepository.GetMemberByPhoneNumber() : %v", err)
		}
		return nil, nil
	}
	return mmb, nil
}

func (mmbRepo *memberRepository) GetMemberByIdMember(idmem string) (*model.Member, error) {
	qry := utils.GET_MEMBER_BY_IDMEMBER
	mmb := &model.Member{}
	err := mmbRepo.db.QueryRow(qry, idmem).Scan(&mmb.Id, &mmb.Name, &mmb.PhoneNo, &mmb.NoIdentity, &mmb.ImagePath, &mmb.LoanStatus, &mmb.CreatedAt, &mmb.UpdatedAt, &mmb.CreatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("error on memberRepository.GetMemberByIdMember() : %v", err)
		}
		return nil, nil
	}
	return mmb, nil
}

func (mmbRepo *memberRepository) GetAllMember() ([]*model.Member, error) {
	qry := utils.GET_ALL_MEMBER
	var arrMember []*model.Member
	rows, err := mmbRepo.db.Query(qry)
	if err != nil {
		return nil, fmt.Errorf("error on memberRepository.GetAllMember() : %v", err)
	}

	for rows.Next() {
		mmb := &model.Member{}
		rows.Scan(&mmb.Id, &mmb.Name, &mmb.PhoneNo, &mmb.NoIdentity, &mmb.ImagePath, &mmb.LoanStatus, &mmb.CreatedAt, &mmb.UpdatedAt, &mmb.CreatedBy)
		arrMember = append(arrMember, mmb)
	}
	return arrMember, nil

}

func (mmbRepo *memberRepository) GetMemberById(id string) (*model.Member, error) {
	qry := utils.GET_MEMBER_BY_ID
	mmb := &model.Member{}
	err := mmbRepo.db.QueryRow(qry, id).Scan(&mmb.Id, &mmb.Name, &mmb.PhoneNo, &mmb.NoIdentity, &mmb.ImagePath, &mmb.LoanStatus, &mmb.CreatedAt, &mmb.UpdatedAt, &mmb.CreatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("error on memberRepository.GetMemberById() : %v", err)
		}
		return nil, nil
	}
	return mmb, nil
}

func (mmbRepo *memberRepository) InsertMember(mmb *model.Member) error {
	qry := utils.INSERT_MEMBER
	_, err := mmbRepo.db.Exec(qry, &mmb.Id, &mmb.Name, &mmb.PhoneNo, &mmb.NoIdentity, &mmb.ImagePath, &mmb.LoanStatus, &mmb.CreatedAt, &mmb.UpdatedAt, &mmb.CreatedBy)
	if err != nil {
		return fmt.Errorf("error on memberRepository.InsertMember() : %w", err)
	}
	return nil
}

func (mmbRepo *memberRepository) DeleteMember(mmb *model.Member) error {
	qry := utils.DELETE_MEMBER
	_, err := mmbRepo.db.Exec(qry, mmb.Id) // Menggunakan mmb.ID sebagai id parameter untuk menghapus anggota
	if err != nil {
		return fmt.Errorf("error on memberRepository.DeleteMember: %w", err)
	}
	return nil
}

func (mmbRepo *memberRepository) UpdateMember(mmb *model.Member) error {
	qry := utils.UPDATE_MEMBER
	_, err := mmbRepo.db.Exec(qry, &mmb.Id, &mmb.Name, &mmb.PhoneNo, &mmb.NoIdentity, &mmb.ImagePath, &mmb.LoanStatus, &mmb.CreatedAt, &mmb.UpdatedAt, &mmb.CreatedBy)
	if err != nil {
		return fmt.Errorf("error on memberRepository.UpdateMember : %v", &err)
	}
	return nil
}

func NewMemberRepository(db *sql.DB) MemberRepository {
	return &memberRepository{
		db: db,
	}
}



