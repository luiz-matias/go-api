package route

import (
	"go-api/api/controller"

	"github.com/gin-gonic/gin"
)

func ConfigureProductRoutes(server *gin.Engine, controller *controller.ProductController) {
	server.GET("/products/:id", controller.GetProduct)
	server.GET("/products", controller.GetProducts)
	server.POST("/products", controller.CreateProduct)
}
