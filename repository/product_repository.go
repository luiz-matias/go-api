package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (p *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT * FROM products"
	rows, err := p.connection.Query(query)

	if err != nil {
		fmt.Println("Error while querying products: " + err.Error())
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
			fmt.Println("Error while querying products: " + err.Error())
			return []model.Product{}, err
		}

		productList = append(productList, productObject)
	}
	rows.Close()

	if productList == nil {
		return []model.Product{}, nil
	}

	return productList, nil
}

func (p *ProductRepository) GetProduct(id int) (*model.Product, error) {
	query, err := p.connection.Prepare("SELECT * FROM products WHERE id = $1 LIMIT 1")

	if err != nil {
		fmt.Println("Error while querying product: " + err.Error())
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
			return nil, nil
		}
		fmt.Println("Error while querying product: " + err.Error())
		return nil, err
	}

	return &productObject, nil
}

func (p *ProductRepository) SaveProduct(product *model.Product) (*model.Product, error) {
	query := "INSERT INTO products(name, price, quantity) VALUES ($1, $2, $3) RETURNING id"
	statement, err := p.connection.Prepare(query)

	if err != nil {
		fmt.Println("Error while inserting product: " + err.Error())
		return nil, err
	}

	var id int
	err = statement.QueryRow(product.Name, product.Price, product.Quantity).Scan(&id)

	if err != nil {
		fmt.Println("Error while inserting product: " + err.Error())
		return nil, err
	}

	return p.GetProduct(id)
}
