package userController

import (
	"net/http"

	"example.com/m/v2/database"
	"example.com/m/v2/models"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var newUser *models.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.User
	result := database.DB.Where("email = ?", newUser.Email).First(&existingUser)
	if result.RowsAffected > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		return
	}

	if !govalidator.IsEmail(newUser.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}

	if len(newUser.Password) < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password is required"})
		return
	}

	if !govalidator.StringLength(newUser.Password, "6", "255") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Panjang harus 6"})
		return
	}

	if err := database.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": newUser})
}
