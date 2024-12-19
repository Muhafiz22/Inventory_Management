package models

import (
	"fmt"
)

type Order struct {
	OrderId    int
	Products   []OrderItem
	TotalPrice float32
	IsCredit   bool
}

type OrderItem struct { //to track only the necessary changes and keep the details current
	ProductId int
	Price     float32
	Quantity  int
	P         *Product
}

func (o *Order) InitialiseOrder(nextOrderId int) {

	o.OrderId = nextOrderId
	o.TotalPrice = 0.0
	o.IsCredit = false
	o.Products = []OrderItem{}
}

func (o *Order) AddProduct(p *Product, Quantity int) bool {

	fmt.Println("sell Quantity", Quantity)
	if p.Quantity < Quantity { //checks if enough Product Quantity is available to sell or to be added in order list
		return false
	}
	p.Quantity -= Quantity
	for i, item := range o.Products { //checks whether same Product is already in order and so adds it in existing order details
		if item.ProductId == p.ID {
			o.Products[i].Quantity += Quantity
			o.TotalPrice += p.Price * float32((Quantity))
			return true
		}
	}
	newProduct := OrderItem{ // to add new Product in order list as no already existing Product found in order
		ProductId: p.ID,
		Price:     float32(p.Price),
		Quantity:  Quantity,
		P:         p,
	}
	o.Products = append(o.Products, newProduct)
	o.TotalPrice += p.Price * float32((Quantity))
	return true
}

func (o *Order) CalculateTotal() float32 {

	o.TotalPrice = 0.0
	for _, item := range o.Products {
		o.TotalPrice += float32(item.Price) * float32(item.Quantity)
	}
	return o.TotalPrice
}

func (o *Order) DisplayOrderDetails() {

	fmt.Println("----------Order Details----------")
	fmt.Println("order Id:", o.OrderId)
	for _, item := range o.Products {
		item.P.Put()
		fmt.Println("Ordered Quantity:", item.Quantity)
	}
	fmt.Println("Total Price:", o.TotalPrice)
}
