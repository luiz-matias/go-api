package mocks

import (
	"go-api/domain/model"
	"go-api/domain/repository"
)

type ProductRepositoryMock struct{}

func NewProductRepositoryMock() repository.ProductRepository {
	return &ProductRepositoryMock{}
}

func (p *ProductRepositoryMock) GetProducts() ([]model.Product, error) {
	var productList []model.Product

	for i := 0; i < 10; i++ {
		productList = append(productList, model.Product{
			ID:       1,
			Name:     "test",
			Price:    9.99,
			Quantity: 3,
		})
	}

	return productList, nil
}

func (p *ProductRepositoryMock) GetProduct(id int) (*model.Product, error) {
	return &model.Product{
		ID:       id,
		Name:     "test",
		Price:    9.99,
		Quantity: 3,
	}, nil
}

func (p *ProductRepositoryMock) SaveProduct(product *model.Product) (*model.Product, error) {
	return &model.Product{
		ID:       1,
		Name:     product.Name,
		Price:    product.Price,
		Quantity: product.Quantity,
	}, nil
}

func (p *ProductRepositoryMock) UpdateProduct(id int, product *model.Product) (*model.Product, error) {
	return &model.Product{
		ID:       id,
		Name:     product.Name,
		Price:    product.Price,
		Quantity: product.Quantity,
	}, nil
}

func (p *ProductRepositoryMock) DeleteProduct(id int) error {
	return nil
}
