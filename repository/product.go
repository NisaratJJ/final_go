package repository

import (
	"final_go/model"

	"gorm.io/gorm"
)

// สร้าง interface สำหรับ ProductRepository
type ProductRepository interface {
	SearchProducts(description string, minPrice, maxPrice float64) ([]model.Product, error)
}

// ใช้ struct สำหรับ implement interface
type productRepository struct {
	db *gorm.DB
}

// ฟังก์ชันนี้จะใช้ในการค้นหาสินค้าตามคำอธิบายและช่วงราคา
func (repo *productRepository) SearchProducts(description string, minPrice, maxPrice float64) ([]model.Product, error) {
	var products []model.Product

	// ค้นหาสินค้าที่ตรงกับคำอธิบาย และอยู่ในช่วงราคาที่กำหนด
	err := repo.db.Where("description LIKE ? AND price BETWEEN ? AND ?", "%"+description+"%", minPrice, maxPrice).Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

// ฟังก์ชันนี้ใช้สร้าง instance ของ productRepository และ return กลับเป็น interface
func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}
