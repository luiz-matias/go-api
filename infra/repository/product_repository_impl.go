package repository

import (
	"database/sql"
	"go-api/common"
	"go-api/domain/model"
	"go-api/domain/repository"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) repository.ProductRepository {
	return &ProductRepository{
		connection: connection,
	}
}

func (p *ProductRepository) productExists(id int) (bool, error) {
	query, err := p.connection.Prepare("SELECT count(id) FROM products WHERE id = $1 LIMIT 1")

	if err != nil {
		return false, err
	}

	var productId int
	err = query.QueryRow(id).Scan(&productId)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT * FROM products"
	rows, err := p.connection.Query(query)

	if err != nil {
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObject model.Product

	for rows.Next() {
		err = rows.Scan(
			&productObject.ID,
			&productObject.Name,
			&productObject.Price,
			&productObject.Quantity,
		)
		if err != nil {
			return []model.Product{}, err
		}

		productList = append(productList, productObject)
	}
	err = rows.Close()

	if err != nil {
		return []model.Product{}, err
	}

	if productList == nil {
		return []model.Product{}, nil
	}

	return productList, nil
}

func (p *ProductRepository) GetProduct(id int) (*model.Product, error) {
	query, err := p.connection.Prepare("SELECT id, name, price, quantity FROM products WHERE id = $1 LIMIT 1")

	if err != nil {
		return nil, err
	}

	var productObject model.Product
	err = query.QueryRow(id).Scan(
		&productObject.ID,
		&productObject.Name,
		&productObject.Price,
		&productObject.Quantity,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, common.ErrResourceNotFound
		}
		return nil, err
	}

	return &productObject, nil
}

func (p *ProductRepository) SaveProduct(product *model.Product) (*model.Product, error) {
	query := "INSERT INTO products(name, price, quantity) VALUES ($1, $2, $3) RETURNING id"
	statement, err := p.connection.Prepare(query)

	if err != nil {
		return nil, err
	}

	var id int
	err = statement.QueryRow(product.Name, product.Price, product.Quantity).Scan(&id)

	if err != nil {
		return nil, err
	}

	return p.GetProduct(id)
}

func (p *ProductRepository) UpdateProduct(id int, product *model.Product) (*model.Product, error) {
	productExists, err := p.productExists(id)

	if err != nil {
		return nil, err
	}

	if !productExists {
		return nil, common.ErrResourceNotFound
	}

	query := "UPDATE products SET name = $1, price = $2, quantity = $3 WHERE id = $4 RETURNING id"
	statement, err := p.connection.Prepare(query)

	if err != nil {
		return nil, err
	}

	_, err = statement.Exec(product.Name, product.Price, product.Quantity, id)

	if err != nil {
		return nil, err
	}

	return p.GetProduct(id)
}

func (p *ProductRepository) DeleteProduct(id int) error {
	productExists, err := p.productExists(id)

	if err != nil {
		return err
	}

	if !productExists {
		return common.ErrResourceNotFound
	}

	query := "DELETE FROM products WHERE id = $1"
	statement, err := p.connection.Prepare(query)

	if err != nil {
		return err
	}

	_, err = statement.Exec(id)

	if err != nil {
		return err
	}

	return nil
}
