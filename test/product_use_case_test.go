package use_case

import (
	"go-api/domain/model"
	"go-api/domain/use_case"
	"go-api/infra/mocks"
	"testing"
)

func TestGetProductsUseCase(t *testing.T) {
	repository := mocks.NewProductRepositoryMock()
	usecase := use_case.NewProductUseCase(repository)

	products, err := usecase.GetProducts()

	if err != nil {
		t.Errorf("Error thrown when getting products: %s", err)
	}

	if len(products) != 10 {
		t.Errorf("Expected %d, got %d products", 10, len(products))
	}
}

func TestGetProductUseCase(t *testing.T) {
	repository := mocks.NewProductRepositoryMock()
	usecase := use_case.NewProductUseCase(repository)

	product, err := usecase.GetProductByID(1)

	if err != nil {
		t.Errorf("Error thrown when getting product: %s", err)
	}

	if product.ID != 2 {
		t.Errorf("Expected product id == %d, got %d", 2, product.ID)
	}
}

func TestSaveProductUseCase(t *testing.T) {
	repository := mocks.NewProductRepositoryMock()
	usecase := use_case.NewProductUseCase(repository)

	product, err := usecase.CreateProduct(model.Product{
		Name:     "testing",
		Price:    1.99,
		Quantity: 5,
	})

	if err != nil {
		t.Errorf("Error thrown when saving product: %s", err)
	}

	if product.ID != 1 {
		t.Errorf("Expected product id == %d, got %d", 1, product.ID)
	}

	if product.Name != "testing" {
		t.Errorf("Expected product name == %s, got %s", "testing", product.Name)
	}

	if product.Price != 1.99 {
		t.Errorf("Expected product price == %f, got %f", 1.99, product.Price)
	}

	if product.Quantity != 5 {
		t.Errorf("Expected product quantity == %d, got %d", 5, product.Quantity)
	}
}

func TestUpdateProductUseCase(t *testing.T) {
	repository := mocks.NewProductRepositoryMock()
	usecase := use_case.NewProductUseCase(repository)

	product, err := usecase.UpdateProduct(123, model.Product{
		Name:     "testing",
		Price:    1.99,
		Quantity: 5,
	})

	if err != nil {
		t.Errorf("Error thrown when updating product: %s", err)
	}

	if product.ID != 123 {
		t.Errorf("Expected product id == %d, got %d", 123, product.ID)
	}

	if product.Name != "testing" {
		t.Errorf("Expected product name == %s, got %s", "testing", product.Name)
	}

	if product.Price != 1.99 {
		t.Errorf("Expected product price == %f, got %f", 1.99, product.Price)
	}

	if product.Quantity != 5 {
		t.Errorf("Expected product quantity == %d, got %d", 5, product.Quantity)
	}
}

func TestDeleteProductUseCase(t *testing.T) {
	repository := mocks.NewProductRepositoryMock()
	usecase := use_case.NewProductUseCase(repository)

	err := usecase.DeleteProduct(123)

	if err != nil {
		t.Errorf("Error thrown when deleting product: %s", err)
	}
}
