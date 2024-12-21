package tests_test

import (
	"inventory_management/backend/models"
	"inventory_management/backend/services"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

/*
TO-DO
1.first make arrangements for calling get function
*/

// embedding models.Product struct members
// type MockProduct struct {
// 	models.Product
// 	MockGetFunc func(products []models.Product)
// }

// overriding Get function
// func (mp *MockProduct) Get(products []models.Product) {

// 	if mp.MockGetFunc != nil {
// 		mp.MockGetFunc(products)
// 	}
// }

func TestAddProductWithDynamicInput(t *testing.T) {
	// Prepare initial products
	initialProducts := []models.Product{
		{ID: 1, Name: "Product A", Price: 100.0, Quantity: 10},
		{ID: 2, Name: "Product B", Price: 200.0, Quantity: 20},
	}

	// Create a pipe to simulate step-by-step input
	reader, writer, _ := os.Pipe()

	// Save the original os.Stdin
	originalStdin := os.Stdin
	defer func() { os.Stdin = originalStdin }() // Restore os.Stdin after the test

	// Replace os.Stdin with the reader end of the pipe
	os.Stdin = reader

	// Use a goroutine to write inputs dynamically
	go func() {
		time.Sleep(100 * time.Millisecond) // Allow time for the Get function to request input
		writer.WriteString("3\n")          // Enter Product ID
		time.Sleep(100 * time.Millisecond)
		writer.WriteString("Product C\n") // Enter Product Name
		time.Sleep(100 * time.Millisecond)
		writer.WriteString("300.0\n") // Enter Product Price
		time.Sleep(100 * time.Millisecond)
		writer.WriteString("30\n") // Enter Product Quantity
		writer.Close()             // Close the writer after providing all inputs
	}()

	// Create a ProductService
	ps := &services.ProductService{Product: &models.Product{}}

	// Call the AddProduct function
	updatedProducts, err := ps.AddProduct(initialProducts)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, 3, len(updatedProducts)) // Verify new product is added

	// Check the details of the new product
	newProduct := updatedProducts[2]
	assert.Equal(t, 3, newProduct.ID)
	assert.Equal(t, "Product C", newProduct.Name)
	assert.Equal(t, 300.0, newProduct.Price)
	assert.Equal(t, 30, newProduct.Quantity)
}

// func TestAddProduct(t *testing.T) {

// 	initialProduct := []models.Product{
// 		{ID: 1, Name: "Product A", Price: 100.0, Quantity: 10},
// 		{ID: 2, Name: "Product B", Price: 200.0, Quantity: 20},
// 	}

// 	mockProduct := &MockProduct{}
// 	mockProduct.MockGetFunc = func(products []models.Product) {
// 		// Simulate new product details entered by the user
// 		mockProduct.ID = 3
// 		mockProduct.Name = "Product C"
// 		mockProduct.Price = 300.0
// 		mockProduct.Quantity = 30
// 	}

// 	productService := &services.ProductService{Product: &models.Product{}}
// 	products, err := productService.AddProduct(initialProduct)

// 	assert.NoError(t, err)
// 	assert.Len(t, products, 3)
// 	assert.Equal(t, 2, products[1].ID)
//fmt.Println(products[2].ID) new product is not being added in inventory
//assert.Equal(t, "Product C", products[2].Name)
// assert.Equal(t, 300.0, products[2].Price)
// assert.Equal(t, 30, products[2].Quantity)

// // Update the mock behavior to simulate updating an existing product
// mockProduct.MockGetFunc = func(products []models.Product) {
// 	// Simulate finding an existing product and increasing quantity
// 	mockProduct.ID = 1
// 	mockProduct.Quantity = 15
// }

// // Test updating the quantity of an existing product
// products, err = productService.AddProduct(products)
// assert.NoError(t, err)
// assert.Len(t, products, 3)
// assert.Equal(t, 25, products[0].Quantity)
