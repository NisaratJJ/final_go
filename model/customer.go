package model

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	CustomerID  uint      `gorm:"primaryKey;autoIncrement"`
	FirstName   string    `gorm:"type:varchar(255);not null"`
	LastName    string    `gorm:"type:varchar(255);not null"`
	Email       string    `gorm:"type:varchar(255);unique;not null"`
	PhoneNumber string    `gorm:"type:varchar(20)"`
	Address     string    `gorm:"type:varchar(255)"`
	Password    string    `gorm:"type:varchar(255);not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

func MigrateCustomer(db *gorm.DB) {
	db.AutoMigrate(&Customer{})
}
