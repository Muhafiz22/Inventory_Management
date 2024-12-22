package services

import (
	"fmt"
	"inventory_management/backend/models"
)

type ProductService struct {
	Product *models.Product
}

func (ps *ProductService) AddExistingProductQuantity(existingProduct *models.Product, AdditionalQuantity int) {

	existingProduct.Quantity += AdditionalQuantity
	ps.Product.ID = existingProduct.ID
	ps.Product.Name = existingProduct.Name
	ps.Product.Price = existingProduct.Price
}

func (ps *ProductService) AddProduct(products *[]models.Product, newProduct *models.Product) []models.Product {

	*products = append(*products, *newProduct)
	return *products
}

func (ps *ProductService) handleAddProduct(Products *[]models.Product) {

	p := &models.Product{}

	fmt.Println("Enter product ID: ")
	if _, err := fmt.Scan(&p.ID); err != nil {
		fmt.Println("Invalid input, Please Enter valid Product ID")
	}

	exists, existingProduct := p.CheckExistingProduct(*Products, p.ID)

	if exists {

		fmt.Println("Product already exists")
		fmt.Println("Enter Additional Quantity")

		var AdditionalQuantity int
		fmt.Scan(&AdditionalQuantity)

		ps.AddExistingProductQuantity(existingProduct, AdditionalQuantity)
	} else {

		p.Get()
		ps.AddProduct(Products, p)
	}
}

func (ps *ProductService) displayProducts(products []models.Product) {

	fmt.Println("\nProducts in stock are:")
	for _, item := range products {
		item.Put()
		fmt.Println("Product quantity:", item.Quantity)
		fmt.Println()
	}
}

func (ps *ProductService) searchProduct(products []models.Product) {

	fmt.Println("Enter the product ID to search")
	var key int16
	fmt.Scan(&key)
	found := false
	for _, item := range products {
		if int16(item.ID) == key {
			found = true
			fmt.Println("Product found:", "Id:", item.ID, " Name:", item.Name, " Quantity:", item.Quantity)
			break
		}
	}
	if !found {
		fmt.Println("Out of stock!!")
	}
}
