package repository

import (
	"database/sql"
	"easylib-go/model"
	"easylib-go/utils"
	"fmt"
)

type UserRepository interface {
	GetUserByUsername(string) (*model.User, error)
	AddUser(*model.User) error
	UpdateUser(*model.User) error
	DeleteUser(string) error
}

type userRepository struct {
	db *sql.DB
}

func (usrRepo *userRepository) GetUserByUsername(username string) (*model.User, error) {
	qry := utils.SELECT_USER_BY_USERNAME

	usr := &model.User{}
	err := usrRepo.db.QueryRow(qry, username).Scan(&usr.Id, &usr.Username, &usr.Password, &usr.CreatedAt, &usr.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on userRepository.GetUserByUsername() : %w", err)
	}
	return usr, nil
}

func (usrRepo *userRepository) AddUser(usr *model.User) error {
	qry := utils.INSERT_USER

	_, err := usrRepo.db.Exec(qry, &usr.Id, &usr.Username, &usr.Password, &usr.CreatedAt, &usr.UpdatedAt)
	if err != nil {
		return fmt.Errorf("error on userRepository.AddUser() : %w", err)
	}
	return nil
}

func (usrRepo *userRepository) UpdateUser(usr *model.User) error {
	qry := utils.UPDATE_USER

	_, err := usrRepo.db.Exec(qry, &usr.Username, &usr.Password, &usr.CreatedAt, &usr.UpdatedAt, &usr.Id)
	if err != nil {
		return fmt.Errorf("error on userRepository.UpdateUser() : %w", err)
	}
	return nil
}

func (lprdctRepo *userRepository) DeleteUser(username string) error {
	qry := utils.DELETE_USER
	_, err := lprdctRepo.db.Exec(qry, username)
	if err != nil {
		return fmt.Errorf("error on userRepository.DeleteUser() : %v", err)
	}
	return nil
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}