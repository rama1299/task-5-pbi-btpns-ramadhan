package userRouter

import (
	"example.com/m/v2/controllers/userController"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) {
	userGroup := r.Group("/user")

	userGroup.POST("/register", userController.UserRegister)
	userGroup.POST("/login", userController.UserLogin)
	userGroup.PUT("/:userId", userController.UserUpdate)
	userGroup.DELETE("/:userId", userController.UserDelete)
}
