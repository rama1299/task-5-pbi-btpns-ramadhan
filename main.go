package main

import (
	_ "example.com/m/v2/database"
	"example.com/m/v2/router/userRouter"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	userRouter.UserRouter(r)

	r.Run(":8080")
}
