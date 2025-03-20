package controller

import (
	"go-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	//UseCase there
}

func NewProductController() ProductController {
	return ProductController{}
}

func (p *ProductController) GetProducts(ctx *gin.Context) {
	products := []model.Product{
		{ID: 1, Name: "Laptop", Price: 5000, Quantity: 10},
		{ID: 1, Name: "Mouse", Price: 2000, Quantity: 15},
		{ID: 1, Name: "Headphone", Price: 3000, Quantity: 20},
	}

	ctx.JSON(http.StatusOK, products)
}
