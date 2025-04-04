package route

import (
	"go-api/api/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ConfigureHealthRoutes(server *gin.Engine) {
	server.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "UP",
		})
	})
}

func ConfigureProductRoutes(server *gin.Engine, controller *controller.ProductController) {
	server.GET("/products/:id", controller.GetProduct)
	server.GET("/products", controller.GetProducts)
	server.POST("/products", controller.CreateProduct)
}
