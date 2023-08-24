package userController

import (
	"net/http"
	"time"

	"example.com/m/v2/app"
	"example.com/m/v2/database"
	"example.com/m/v2/helpers"
	"example.com/m/v2/models"
	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var newUser *models.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !helpers.Required(newUser.Username) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
		return
	}

	if !helpers.IsEmail(newUser.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}

	if !helpers.Required(newUser.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password is required"})
		return
	}

	if !helpers.MinlengthPassword(newUser.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be min 6 characters and max 15"})
		return
	}

	var existingUser models.User
	result := database.DB.Where("email = ?", newUser.Email).First(&existingUser)
	if result.RowsAffected > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		return
	}

	hashPassword, err := app.HashPassword(newUser.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	newUser.Password = hashPassword

	if err := database.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Create user successful"})
}

func UserLogin(c *gin.Context) {
	var loginData *app.UserLogin

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !helpers.Required(loginData.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data incomplete"})
		return
	}

	if !helpers.Required(loginData.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data incomplete"})
		return
	}

	if !helpers.IsEmail(loginData.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}

	var existingUser models.User
	result := database.DB.Where("email = ?", loginData.Email).First(&existingUser)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if !app.CheckPassword(loginData.Password, existingUser.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := app.JsonWebToken(int(existingUser.ID), existingUser.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	cookie := &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour * 24),
	}

	http.SetCookie(c.Writer, cookie)

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
func UserUpdate(c *gin.Context) {

}

func UserDelete(c *gin.Context) {

}
