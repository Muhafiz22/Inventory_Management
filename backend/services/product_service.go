package services

import (
	"fmt"
	"inventory_management/backend/models"
)

type ProductService struct {
	Product *models.Product
}

func (ps *ProductService) AddProduct(products []models.Product) ([]models.Product, error) {

	p := &models.Product{}

	p.Get(products)

	existingProduct := models.FindProductById(products, p.ID)
	if existingProduct == nil {
		products = append(products, *p)
	}
	return products, nil
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
