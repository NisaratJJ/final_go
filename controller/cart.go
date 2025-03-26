package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewCart(router *gin.Engine, db *gorm.DB) {
	dbCus = db
	user := router.Group("/cart")
	{
		user.GET("/", getAllCart)

	}
}

func getAllCart(ctx *gin.Context) {

}
