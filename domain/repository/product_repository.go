package repository

import "go-api/domain/model"

type ProductRepository interface {
	GetProducts() ([]model.Product, error)
	GetProduct(id int) (*model.Product, error)
	SaveProduct(product *model.Product) (*model.Product, error)
	UpdateProduct(id int, product *model.Product) (*model.Product, error)
	DeleteProduct(id int) error
}
