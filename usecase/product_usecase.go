package usecase

import (
	"go-api/model"
	"go-api/repository"
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
	products, err := productUseCase.repository.GetProducts()

	if err != nil {
		return []model.Product{}, err
	}

	return products, nil
}

func (productUseCase *ProductUseCase) GetProductByID(productId int) (*model.Product, error) {
	product, err := productUseCase.repository.GetProduct(productId)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (productUseCase *ProductUseCase) CreateProduct(product model.Product) (*model.Product, error) {
	savedProduct, err := productUseCase.repository.SaveProduct(&product)

	if err != nil {
		return nil, err
	}

	return savedProduct, nil
}
