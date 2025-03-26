package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func StartServer() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", viper.GetString("mysql.dsn"))
	dsn := viper.GetString("mysql.dsn")
	dialector := mysql.Open(dsn)
	db, err := gorm.Open(dialector)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection Success")

	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "message API Is Working....")
	})

	NewCustomer(router, db)
	NewCart(router, db)
	NewProduct(router, db)

	router.Run()
}
