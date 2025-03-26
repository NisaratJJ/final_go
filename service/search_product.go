package service

import (
	"final_go/model"
	"final_go/repository"

	"gorm.io/gorm"
)

// กำหนด interface สำหรับบริการค้นหาสินค้า
type SearchProductsService interface {
	SearchProducts(description string, minPrice float64, maxPrice float64) ([]model.Product, error)
}

type productDB struct {
	db *gorm.DB
}

// การ implement ฟังก์ชัน SearchProducts จาก interface
func (p *productDB) SearchProducts(description string, minPrice float64, maxPrice float64) ([]model.Product, error) {
	// ใช้ repository ในการค้นหาสินค้า
	productRepo := repository.NewProductRepository(p.db)
	products, err := productRepo.SearchProducts(description, minPrice, maxPrice)
	if err != nil {
		return nil, err
	}
	return products, nil
}

// ฟังก์ชันนี้สร้าง instance ของ productDB ที่ implement SearchProductsService
func NewProductService(db *gorm.DB) SearchProductsService {
	return &productDB{db: db}
}
