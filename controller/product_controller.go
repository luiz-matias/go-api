package controller

import (
	"go-api/dto"
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productUseCase usecase.ProductUseCase
}

func NewProductController(usecase usecase.ProductUseCase) ProductController {
	return ProductController{
		productUseCase: usecase,
	}
}

func (p *ProductController) GetProducts(ctx *gin.Context) {
	products, err := p.productUseCase.GetProducts()

	if err != nil {
		response := dto.ErrorResponse{
			Message:    "Internal server error",
			Stacktrace: err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *ProductController) CreateProduct(ctx *gin.Context) {

	var productModel model.Product
	err := ctx.BindJSON(&productModel)

	if err != nil {
		response := dto.ErrorResponse{
			Message:    "Internal server error",
			Stacktrace: err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	product, err := p.productUseCase.CreateProduct(productModel)

	if err != nil {
		response := dto.ErrorResponse{
			Message:    "Internal server error",
			Stacktrace: err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Product created successfully",
		"product": product,
	})
}

func (p *ProductController) GetProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		response := dto.MessageResponse{
			Message: "Invalid product id",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)

	if err != nil {
		response := dto.MessageResponse{
			Message: "Invalid product id",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.productUseCase.GetProductByID(productId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message":    "Internal server error",
			"stacktrace": err.Error(),
		})
		return
	}

	if product == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Product not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, product)
}
