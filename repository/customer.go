package repository

import (
	"final_go/model"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	GetCustomerByEmail(email string) (*model.Customer, error)
	GetAllCustomer() (*[]model.Customer, error) // Removed email parameter
	UpdateAddress(id int, customer model.Customer) (int, error)
	UpdatePassword(id int, newPassword string) (int, error)
	GetCustomerById(id int) (*model.Customer, error)
}

type customerDB struct {
	db *gorm.DB
}

func (c customerDB) GetCustomerById(id int) (*model.Customer, error) {
	var customer model.Customer
	// ใช้คำสั่ง Where เพื่อค้นหาผู้ใช้ที่มี customer_id ตรงกับ id
	result := c.db.Where("customer_id = ?", id).First(&customer)
	if result.Error != nil {
		return nil, result.Error
	}
	return &customer, nil
}

func (s customerDB) UpdatePassword(id int, newPassword string) (int, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return -1, err
	}
	result := s.db.Model(&model.Customer{}).Where("customer_id = ?", id).Update("password", string(hashedPassword))

	if result.Error != nil {
		return -1, result.Error
	}

	// ถ้าอัปเดตสำเร็จ
	return id, nil
}

func (s customerDB) UpdateAddress(id int, user model.Customer) (int, error) {
	result := s.db.Model(model.Customer{}).Where("customer_id = ?", id).Updates(user)
	if result.Error != nil {
		return -1, nil
	} else {

		return id, nil
	}

}
func (c customerDB) GetAllCustomer() (*[]model.Customer, error) {
	customers := []model.Customer{}
	result := c.db.Find(&customers)
	if result.Error != nil {
		return nil, result.Error
	}
	return &customers, nil
}

func (c customerDB) GetCustomerByEmail(email string) (*model.Customer, error) {
	customers := model.Customer{}
	if err := c.db.Where("email = ?", email).First(&customers).Error; err != nil {
		return nil, err
	}
	fmt.Print(00000)
	return &customers, nil
}

func NewCustomerRepository(db *gorm.DB) customerDB {
	return customerDB{db: db}
}
