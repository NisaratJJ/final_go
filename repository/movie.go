package repository

import (
	"final_go/model"

	"gorm.io/gorm"
)

type MovieRepository interface {
	GetAll() (*[]model.Country, error)
	// GetByName(name string) (*[]model.Country, error)
}

type movieDB struct {
	db *gorm.DB
}

func (c movieDB) GetAll() (*[]model.Country, error) {
	countries := []model.Country{}
	result := c.db.Find(&countries)
	if result.Error != nil {
		return nil, result.Error
	}
	return &countries, nil
}

func NewMovieRepository(gormdb *gorm.DB) MovieRepository {
	return movieDB{db: gormdb}
}
