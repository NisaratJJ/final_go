package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var dbCus *gorm.DB

func NewCustomer(router *gin.Engine, db *gorm.DB) {
	dbCus = db
	user := router.Group("/customer")
	{
		user.POST("/auth/login", login)

	}
}

func login(ctx *gin.Context) {
	fmt.Print("55555")
	// service := service.NewShowData(dbMenu)
	// data, err := service.GetAllCountries()
	// if err != nil {
	// 	panic(err)
	// }
	// ctx.JSON(http.StatusOK, data)
}
