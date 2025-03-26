package service

import (
	"final_go/model"
	"final_go/repository"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	UpdateAddressUser(id int, user model.Customer) (int, error)
	UpdatePasswordUser(id int, oldPassword string, newPassword string) (int, error)
}

type userData struct {
	db *gorm.DB
}

func (u userData) UpdatePasswordUser(id int, oldPassword string, newPassword string) (int, error) {

	customerRepo := repository.NewCustomerRepository(u.db)
	customer, err := customerRepo.GetCustomerById(id)
	if err != nil {
		return 0, err
	}

	// ตรวจสอบรหัสผ่านเก่ากับรหัสผ่านที่เก็บไว้ในฐานข้อมูล
	if err := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(oldPassword)); err != nil {
		return 0, fmt.Errorf("old password is incorrect") // ถ้ารหัสผ่านเก่าผิด
	}

	// แฮชรหัสผ่านใหม่ก่อนที่จะเก็บลงฐานข้อมูล
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return 0, fmt.Errorf("failed to hash new password: %v", err)
	}

	update := repository.NewCustomerRepository(u.db)
	result, err := update.UpdatePassword(id, string(hashedPassword))
	if err != nil {
		return 0, err
	}

	return result, nil
}

// เพิ่มเมธอด UpdateAddressUser ใน userData
func (u userData) UpdateAddressUser(id int, user model.Customer) (int, error) {
	customerRepo := repository.NewCustomerRepository(u.db)

	// อัปเดตข้อมูลที่อยู่ของลูกค้า
	result, err := customerRepo.UpdateAddress(id, user)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func NewUserService(gormdb *gorm.DB) UserService {
	return userData{db: gormdb}
}
