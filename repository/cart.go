package repository

import (
	"final_go/model"

	"gorm.io/gorm"
)

type CartRepository interface {
	GetAllCart(email string) (*model.CartItem, error)
}

type cartDB struct {
	db *gorm.DB
}
