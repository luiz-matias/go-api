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

func (productUseCase *ProductUseCase) GetProducts() ([]model.Product, error) {
	return productUseCase.repository.GetProducts()
}

func (productUseCase *ProductUseCase) GetProductByID(productId int) (*model.Product, error) {
	return productUseCase.repository.GetProduct(productId)
}

func (productUseCase *ProductUseCase) CreateProduct(product model.Product) (*model.Product, error) {
	return productUseCase.repository.SaveProduct(&product)
}

func (productUseCase *ProductUseCase) UpdateProduct(productId int, product model.Product) (*model.Product, error) {
	return productUseCase.repository.UpdateProduct(productId, &product)
}

func (productUseCase *ProductUseCase) DeleteProduct(productId int) error {
	return productUseCase.repository.DeleteProduct(productId)
}
