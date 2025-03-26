package service

import (
	"final_go/model"
	"final_go/repository"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginService interface {
	Login(email string, password string) (*[]model.Customer, error)
	UserAll() (*[]model.Customer, error)
}
type customerData struct {
	db *gorm.DB
}

func (c customerData) UserAll() (*[]model.Customer, error) {
	customerRepo := repository.NewCustomerRepository(c.db)
	customer, err := customerRepo.GetAllCustomer()
	if err != nil {
		return nil, err
	}
	return customer, nil
}
func (c customerData) Login(email string, password string) (*model.Customer, error) {
	fmt.Print(888)
	customerRepo := repository.NewCustomerRepository(c.db)
	customer, err := customerRepo.GetCustomerByEmail(email)
	if err != nil {
		return nil, err
	}
	fmt.Print(7777)
	// Compare hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(password)); err != nil {
		return nil, err
	}
	fmt.Print(customer)
	return customer, nil
}

func NewLoginService(gormdb *gorm.DB) customerData {
	return customerData{db: gormdb}
}
