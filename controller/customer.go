package controller

import (
	"final_go/dto"
	"final_go/model"
	"final_go/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var dbCus *gorm.DB

func NewCustomer(router *gin.Engine, db *gorm.DB) {
	dbCus = db
	user := router.Group("/customer")
	{
		user.POST("/auth/login", login)
		user.GET("/", getAllUser)
		user.PUT("/update", updateUser)
		user.PUT("/change-password", changePassword)
	}
}
func changePassword(ctx *gin.Context) {
	var changePasswordRequest struct {
		CustomerID  int    `json:"customer_id" binding:"required"`
		OldPassword string `json:"old_password" binding:"required"`
		NewPassword string `json:"new_password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&changePasswordRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	service := service.NewUserService(dbCus)
	result, err := service.UpdatePasswordUser(changePasswordRequest.CustomerID, changePasswordRequest.OldPassword, changePasswordRequest.NewPassword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found or old password is incorrect"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}

func updateUser(ctx *gin.Context) {
	var user model.Customer
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data", "details": err.Error()})
		return
	}

	service := service.NewUserService(dbCus)
	updateResult, err := service.UpdateAddressUser(user.CustomerID, user)

	// ถ้ามีข้อผิดพลาดในการอัปเดต
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user", "details": err.Error()})
		return
	}

	// ถ้าอัปเดตสำเร็จ ส่งผลลัพธ์กลับ
	if updateResult == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found or no changes made"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "updated_rows": updateResult})
}

func getAllUser(ctx *gin.Context) {
	services := service.NewLoginService(dbCus)
	users, err := services.UserAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

// func changPassword

func login(ctx *gin.Context) {
	var loginRequest struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// รับ JSON และตรวจสอบข้อผิดพลาด
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	fmt.Print(3333)
	services := service.NewLoginService(dbCus)
	user, err := services.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}
	fmt.Print(2222)
	//แปลงเป็น DTO
	userDTO := dto.UserDTO{
		CustomerID:  int(user.CustomerID),
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
	}
	fmt.Print(111)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user":    userDTO,
	})
}
