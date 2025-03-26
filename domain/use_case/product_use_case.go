package use_case

import (
	"go-api/domain/model"
	"go-api/domain/repository"
)

type ProductUseCase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repository repository.ProductRepository) ProductUseCase {
	return ProductUseCase{
		repository: repository,
	}
}

func (productService *ProductUseCase) GetProducts() ([]model.Product, error) {
	return productService.repository.GetProducts()
}

func (productService *ProductUseCase) GetProductByID(productId int) (*model.Product, error) {
	return productService.repository.GetProduct(productId)
}

func (productService *ProductUseCase) CreateProduct(product model.Product) (*model.Product, error) {
	return productService.repository.SaveProduct(&product)
}
