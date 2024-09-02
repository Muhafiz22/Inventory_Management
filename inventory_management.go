package main

import (
	"fmt"
)

type product struct {
	id       int
	name     string
	price    float32
	quantity int
}

func (p *product) get(Product []product) {
	fmt.Print("Enter product id:")
	fmt.Scan(&p.id)
	existingProduct := findProductById(Product, p.id)
	if existingProduct != nil {
		fmt.Print("Enter additional quantity:")
		var additionalQuantity int
		fmt.Scan(&additionalQuantity)
		existingProduct.quantity += additionalQuantity
		p.id = existingProduct.id
		p.name = existingProduct.name
		p.price = existingProduct.price
		p.quantity = existingProduct.quantity
	} else {
		fmt.Println("Enter product name:")
		fmt.Scan(&p.name)
		fmt.Println("Enter product price:")
		fmt.Scan(&p.price)
		fmt.Println("Enter product quantity:")
		fmt.Scan(&p.quantity)
	}
}

func (p *product) put() {
	fmt.Println("Product Id:", p.id)
	fmt.Println("Product Name:", p.name)
	fmt.Println("Product Price:", p.price)
	fmt.Println("Product Quantity:", p.quantity)
}

func findProductById(products []product, id int) *product {
	for i := range products {
		if products[i].id == id {
			return &products[i]
		}
	}
	return nil
}

type Order struct {
	OrderId    int
	products   []product
	TotalPrice float64
	isCredit   bool
}

func (o *Order) order() {
	o.TotalPrice = 0.0
	o.isCredit = false
	o.products = []product{}
}

func (o *Order) addProduct(p *product, quantity int) bool {
	if p.quantity < quantity {
		return false
	}
	p.quantity -= quantity
	for i, item := range o.products {
		if item.id == p.id {
			o.products[i].quantity += quantity
			return true
		}
	}
	newProduct := *p
	newProduct.quantity = quantity
	o.products = append(o.products, newProduct)
	return true
}

func (o *Order) calculateTotal() float64 {
	o.TotalPrice = 0.0
	for _, item := range o.products {
		o.TotalPrice += float64(item.price) * float64(item.quantity)
	}
	return o.TotalPrice
}

func (o *Order) displayOrderDetails() {
	fmt.Println("Order Id:", o.OrderId)
	fmt.Println("Products in the order:")
	for _, item := range o.products {
		item.put()
	}
	fmt.Println("Total price:", o.TotalPrice)
}

type Store struct {
	products    []product
	orders      []Order
	nextOrderId int
}

func (s *Store) addProduct() {
	var p product
	p.get(s.products)
	existingProduct := findProductById(s.products, p.id)
	if existingProduct == nil {
		s.products = append(s.products, p)
	}
}

func (s *Store) displayProducts() {
	fmt.Println("\nProducts in stock are:")
	for _, item := range s.products {
		item.put()
		fmt.Println()
	}
}

func (s *Store) searchProduct() {
	fmt.Println("Enter the product id to search")
	var key int16
	fmt.Scan(&key)
	found := false
	for _, item := range s.products {
		if int16(item.id) == key {
			found = true
			fmt.Println("Product found:", "Id:", item.id, " Name:", item.name, " Quantity:", item.quantity)
			break
		}
	}
	if !found {
		fmt.Println("Out of stock!!")
	}
}

func (s *Store) createOrder() {
	o := Order{OrderId: s.nextOrderId}
	s.nextOrderId++
	o.order()
	for {
		var id, quantity int
		fmt.Println("Enter product id and quantity:")
		fmt.Scan(&id, &quantity)
		product := findProductById(s.products, id)
		if product == nil {
			fmt.Println("Product not found!!")
		} else if product.quantity < quantity {
			fmt.Println("Product out of stock!!")
		} else {
			if !o.addProduct(product, quantity) {
				fmt.Println("Product out of stock!!")
			}
		}
		var choice string
		fmt.Print("Want to add more items to your Order? (y/n): ")
		fmt.Scan(&choice)
		if choice == "n" || choice == "N" {
			break
		}
	}
	o.calculateTotal()
	o.displayOrderDetails()
	s.orders = append(s.orders, o)
}

func (s *Store) displayAllOrders() {
	fmt.Println("All Orders:")
	for _, order := range s.orders {
		order.displayOrderDetails()
		fmt.Println()
	}
}

func main() {
	store := Store{}

	for {
		var ch int
		fmt.Print("1. Add Product\n2. Display Products\n3. Search Product\n4. Buy\n5. Display All Orders\n6. Exit\n")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&ch)

		switch ch {
		case 1:
			store.addProduct()
		case 2:
			store.displayProducts()
		case 3:
			store.searchProduct()
		case 4:
			store.createOrder()
		case 5:
			store.displayAllOrders()
		case 6:
			fmt.Println("Exiting the program!!")
			return
		default:
			fmt.Println("Invalid choice!!")
		}
	}
}
