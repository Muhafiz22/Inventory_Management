package models

import (
	"fmt"
)

type Product struct {
	ID       int
	Name     string
	Price    float64
	Quantity int
}

func (p *Product) CheckExistingProduct(Products []Product, ID int) (bool, *Product) {

	existingProduct := FindProductById(Products, ID)

	if existingProduct != nil {
		return true, existingProduct
	} else {
		return false, existingProduct
	}
}

func (p *Product) Get() []Product { //to Get Product details

	fmt.Println("Enter Product ID:")
	fmt.Scan(&p.ID)
	fmt.Println("Enter Product Name:")
	fmt.Scan(&p.Name)
	fmt.Println("Enter Product Price:")
	fmt.Scan(&p.Price)
	fmt.Println("Enter Product Quantity:")
	fmt.Scan(&p.Quantity)

	return []Product{{ID: p.ID, Name: p.Name, Price: p.Price, Quantity: p.Quantity}}
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
