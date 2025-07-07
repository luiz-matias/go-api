package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ConfigureHealthRoutes(server *gin.Engine) {
	server.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "UP",
		})
	})
	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})
}
