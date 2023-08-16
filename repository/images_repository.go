package repository

import (
	"database/sql"
	"easylib-go/model"
	"easylib-go/utils"
	"fmt"
)

type ImagesRepository interface {
	GetImageById(string) (*model.Images, error)
	InsertImage(model.Images) model.Images
	DeleteImage(string) error
}

type imagesRepository struct {
	db *sql.DB
}

func (imgRepo *imagesRepository) GetImageById(id string) (*model.Images, error) {
	qry := utils.GET_IMAGE_BY_ID
	img := &model.Images{}
	err := imgRepo.db.QueryRow(qry, id).Scan(&img.Id, &img.Path, &img.CreatedAt, &img.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on imagesRepository.getImageById() : %v", err)
	}
	return img, nil
}

func (imgRepo *imagesRepository) InsertImage(img model.Images) model.Images {
	qry := utils.INSERT_IMAGE
	_, err := imgRepo.db.Exec(qry, &img.Id, &img.Path, &img.CreatedAt, &img.UpdatedAt)
	if err != nil {
		panic(err)
	}
	return img
}

func (imgRepo *imagesRepository) DeleteImage(id string) error {
	qry := utils.DELETE_IMAGE
	_, err := imgRepo.db.Exec(qry, id)
	if err != nil {
		return fmt.Errorf("error on imagesRepository.DeleteImage : %v", err)
	}
	return nil
}

// func (imgRepo *imagesRepository) GetImageByName(name string) (*model.Image, error) {
// 	qry := utils.GET_SERVICE_BY_NAME

// 	img := &model.Image{}
// 	err := imgRepo.db.QueryRow(qry, name).Scan(&img.Id, &img.Name, &img.Uom, &img.Price)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, nil
// 		}
// 		return nil, fmt.Errorf("error on imagesRepository.GetImageByName() : %w", err)
// 	}
// 	return img, nil
// }

func NewImageRepository(db *sql.DB) ImagesRepository {
	return &imagesRepository{
		db: db,
	}
}



