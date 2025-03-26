package service

import (
	"final_go/model"
	"final_go/repository"
	"fmt"

	"gorm.io/gorm"
)

type ShowDataService interface {
	GetAllCountries() (*[]model.Country, error)
}
type showData struct {
	db *gorm.DB
}

func (s showData) GetAllCountries() (*[]model.Country, error) {
	movieRepo := repository.NewMovieRepository(s.db)
	movies, err := movieRepo.GetAll()

	if err != nil {
		panic(err)
	}

	for _, v := range *movies {
		fmt.Printf("%v", v)
	}
	return movies, err
}

func NewShowData(gormdb *gorm.DB) showData {
	return showData{db: gormdb}
}
