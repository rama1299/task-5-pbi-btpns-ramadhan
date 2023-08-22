package userRouter

import (
	"example.com/m/v2/controllers/userController"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) {
	userGroup := r.Group("/user")

	userGroup.POST("/register", userController.Register)
	userGroup.POST("/login")
	userGroup.PUT("/:userId")
	userGroup.DELETE("/:userId")
}
