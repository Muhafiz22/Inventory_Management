package models

import (
	"fmt"
)

type Product struct {
	ID       int
	Name     string
	Price    float32
	Quantity int
}

func (p *Product) Get(Products []Product) { //to Get Product details
	//It Accepts Products slice to check if the product to be added is already in the inventory's product list

	fmt.Print("Enter Product ID:")
	if _, err := fmt.Scan(&p.ID); err != nil {
		fmt.Println("Invalid input, Please Enter valid ID")
	}

	existingProduct := FindProductById(Products, p.ID) // to find if same Product is already exixting in inventory or not

	if existingProduct != nil { // if same Product is already existing in inventory then add in existing Product details

		fmt.Print("Enter additional Quantity:")
		var QdditionalQuantity int
		fmt.Scan(&QdditionalQuantity)

		existingProduct.Quantity += QdditionalQuantity
		p.ID = existingProduct.ID
		p.Name = existingProduct.Name
		p.Price = existingProduct.Price
		p.Quantity = existingProduct.Quantity
	} else { // adding new Product to inventory

		fmt.Println("Enter Product Name:")
		fmt.Scan(&p.Name)
		fmt.Println("Enter Product Price:")
		fmt.Scan(&p.Price)
		fmt.Println("Enter Product Quantity:")
		fmt.Scan(&p.Quantity)
	}
}

func (p *Product) Put() {

	fmt.Println("Product Id:", p.ID)
	fmt.Println("Product Name:", p.Name)
	fmt.Println("Product Price:", p.Price)
	//fmt.Println("Product Quantity:", p.Quantity)
}

func FindProductById(products []Product, ID int) *Product {

	for i := range products {
		if products[i].ID == ID {
			return &products[i]
		}
	}
	return nil
}
