package tests_test

import (
	"inventory_management/backend/models"
	"inventory_management/backend/services"
	"testing"
)

func TestAddExistingProductQuantity(t *testing.T) {
	// Arrange
	existingProduct := &models.Product{
		ID:       1,
		Name:     "Test Product",
		Price:    10.0,
		Quantity: 5,
	}
	ps := &services.ProductService{Product: &models.Product{}}
	additionalQuantity := 10

	// Act
	ps.AddExistingProductQuantity(existingProduct, additionalQuantity)

	// Assert
	expectedQuantity := 15
	if existingProduct.Quantity != expectedQuantity {
		t.Errorf("expected quantity %d, but got %d", expectedQuantity, existingProduct.Quantity)
	}

	if ps.Product.ID != existingProduct.ID || ps.Product.Name != existingProduct.Name || ps.Product.Price != existingProduct.Price {
		t.Errorf("expected product details to match, but they do not")
	}
}

func TestAddProduct(t *testing.T) {
	// Arrange
	products := []models.Product{
		{ID: 1, Name: "Existing Product", Price: 10.0, Quantity: 5},
	}
	newProduct := &models.Product{
		ID:       2,
		Name:     "New Product",
		Price:    20.0,
		Quantity: 10,
	}
	ps := &services.ProductService{}

	// Act
	ps.AddProduct(&products, newProduct)

	// Assert
	if len(products) != 2 {
		t.Errorf("expected products length %d, but got %d", 2, len(products))
	}

	lastProduct := products[len(products)-1]
	if lastProduct.ID != newProduct.ID || lastProduct.Name != newProduct.Name || lastProduct.Price != newProduct.Price || lastProduct.Quantity != newProduct.Quantity {
		t.Errorf("expected last product to match new product, but they do not")
	}
}
