package services

import (
	"fmt"
	"inventory_management/backend/models"
)

type OrderService struct {
	Order *models.Order
}

type Bill struct {
	products    []models.Product
	orders      []models.Order
	nextOrderId int
}

var o *models.Order

func (s *Bill) createOrder() {

	s.nextOrderId++
	o := models.Order{OrderId: s.nextOrderId}
	o.InitialiseOrder(s.nextOrderId)
	for {
		var id, quantity int
		fmt.Println("Enter product id and quantity:")
		fmt.Scan(&id, &quantity)

		product := models.FindProductById(s.products, id)
		if product == nil {
			fmt.Println("Product not found!!")
		} else if product.Quantity < quantity {
			fmt.Println("Product out of stock!!")
		} else {
			if !o.AddProduct(product, quantity) {
				fmt.Println("Product out of stock!!")
			}
		}

		var choice string
		fmt.Print("Want to add more items to your order? (y/n): ")
		fmt.Scan(&choice)
		if choice == "n" || choice == "N" {
			break
		}
	}
	o.CalculateTotal()
	o.DisplayOrderDetails()
	s.orders = append(s.orders, o)
}

func (s *Bill) displayAllOrders() {

	fmt.Println("All Orders:")
	for _, order := range s.orders {
		order.DisplayOrderDetails()
		fmt.Println()
	}
}
