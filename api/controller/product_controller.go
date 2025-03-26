package controller

import (
	"go-api/common"
	"go-api/domain/model"
	"go-api/domain/use_case"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	useCase use_case.ProductUseCase
}

func NewProductController(useCase use_case.ProductUseCase) ProductController {
	return ProductController{
		useCase: useCase,
	}
}

func (p *ProductController) GetProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		common.HandleError(ctx, common.ErrBadRequest)
		return
	}

	productId, err := strconv.Atoi(id)

	if err != nil {
		common.HandleError(ctx, common.ErrBadRequest)
		return
	}

	product, err := p.useCase.GetProductByID(productId)

	if err != nil {
		common.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (p *ProductController) GetProducts(ctx *gin.Context) {
	products, err := p.useCase.GetProducts()

	if err != nil {
		common.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *ProductController) CreateProduct(ctx *gin.Context) {
	var productModel model.Product
	err := ctx.BindJSON(&productModel)

	if err != nil {
		common.HandleError(ctx, err)
		return
	}

	product, err := p.useCase.CreateProduct(productModel)

	if err != nil {
		common.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Product created successfully",
		"product": product,
	})
}

func (p *ProductController) UpdateProduct(ctx *gin.Context) {
	var productModel model.Product
	err := ctx.BindJSON(&productModel)

	if err != nil {
		common.HandleError(ctx, err)
		return
	}

	product, err := p.useCase.CreateProduct(productModel)

	if err != nil {
		common.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Product created successfully",
		"product": product,
	})
}
