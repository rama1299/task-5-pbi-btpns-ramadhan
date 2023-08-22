package photosRouter

import (
	"github.com/gin-gonic/gin"
)

func PhotosRouter(r *gin.Engine) {
	photosGroup := r.Group("/photos")

	photosGroup.GET("/")
	photosGroup.GET("/:photoId")
	photosGroup.POST("/")
	photosGroup.DELETE("/:photoId")
}
