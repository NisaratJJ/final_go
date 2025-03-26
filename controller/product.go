package controller

import (
	"final_go/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewProduct(router *gin.Engine, db *gorm.DB) {
	dbCus = db
	product := router.Group("/product")
	{
		product.GET("/search", searchProducts)
	}
}

// ฟังก์ชันนี้จะเป็น API ที่ใช้ค้นหาสินค้า
func searchProducts(ctx *gin.Context) {
	// ดึงข้อมูลจาก query parameters
	description := ctx.Query("description")
	// description := ctx.DefaultQuery("description", "")
	// minPriceStr := ctx.DefaultQuery("min_price", "0")
	minPriceStr := ctx.Query("min_price")
	maxPriceStr := ctx.Query("max_price")
	// maxPriceStr := ctx.DefaultQuery("max_price", "1000000")
	// แปลงช่วงราคาจาก string เป็น float64
	minPrice, err := strconv.ParseFloat(minPriceStr, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid min_price"})
		return
	}

	maxPrice, err := strconv.ParseFloat(maxPriceStr, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid max_price"})
		return
	}

	// เรียกใช้ service เพื่อค้นหาสินค้า
	service := service.NewProductService(dbCus)
	products, err := service.SearchProducts(description, minPrice, maxPrice)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ส่งข้อมูลสินค้าที่ค้นหากลับ
	ctx.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}
