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

func (p *product) get(Product []product) { //to get product details

	fmt.Print("Enter product id:")
	fmt.Scan(&p.id)

	existingProduct := findProductById(Product, p.id) // to find if same product is already exixting in inventory or not

	if existingProduct != nil { // if same product is already existing in inventory then add in existing product details

		fmt.Print("Enter additional quantity:")
		var additionalQuantity int
		fmt.Scan(&additionalQuantity)

		existingProduct.quantity += additionalQuantity
		p.id = existingProduct.id
		p.name = existingProduct.name
		p.price = existingProduct.price
		p.quantity = existingProduct.quantity
	} else { // adding new product to inventory

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
	//fmt.Println("Product Quantity:", p.quantity)
}

func findProductById(products []product, id int) *product {

	for i := range products {
		if products[i].id == id {
			return &products[i]
		}
	}
	return nil
}

type order struct {
	orderId    int
	products   []orderItem
	totalPrice float32
	isCredit   bool
}

type orderItem struct { //to track only the necessary changes and keep the details current
	productId int
	price     float32
	quantity  int
	p         *product
}

func (o *order) initialiseOrder(nextOrderId int) {

	o.orderId = nextOrderId
	o.totalPrice = 0.0
	o.isCredit = false
	o.products = []orderItem{}
}

func (o *order) addProduct(p *product, quantity int) bool {

	fmt.Println("sell quantity", quantity)
	if p.quantity < quantity { //checks if enough product quantity is available to sell or to be added in order list
		return false
	}
	p.quantity -= quantity
	for i, item := range o.products { //checks whether same product is already in order and so adds it in existing order details
		if item.productId == p.id {
			o.products[i].quantity += quantity
			o.totalPrice += p.price * float32((quantity))
			return true
		}
	}
	newProduct := orderItem{ // to add new product in order list as no already existing product found in order
		productId: p.id,
		price:     float32(p.price),
		quantity:  quantity,
		p:         p,
	}
	o.products = append(o.products, newProduct)
	o.totalPrice += p.price * float32((quantity))
	return true
}

func (o *order) calculateTotal() float32 {

	o.totalPrice = 0.0
	for _, item := range o.products {
		o.totalPrice += float32(item.price) * float32(item.quantity)
	}
	return o.totalPrice
}

func (o *order) displayOrderDetails() {

	fmt.Println("----------Order Details----------")
	fmt.Println("order Id:", o.orderId)
	for _, item := range o.products {
		item.p.put()
		fmt.Println("Ordered quantity:", item.quantity)
	}
	fmt.Println("Total price:", o.totalPrice)
}

type store struct {
	products    []product
	orders      []order
	nextOrderId int
}

func (s *store) addProduct() {

	var p product
	p.get(s.products)

	existingProduct := findProductById(s.products, p.id)
	if existingProduct == nil {
		s.products = append(s.products, p)
	}
}

func (s *store) displayProducts() {

	fmt.Println("\nProducts in stock are:")
	for _, item := range s.products {
		item.put()
		fmt.Println("Product quantity:", item.quantity)
		fmt.Println()
	}
}

func (s *store) searchProduct() {

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

func (s *store) createOrder() {

	s.nextOrderId++
	o := order{orderId: s.nextOrderId}
	o.initialiseOrder(s.nextOrderId)
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
		fmt.Print("Want to add more items to your order? (y/n): ")
		fmt.Scan(&choice)
		if choice == "n" || choice == "N" {
			break
		}
	}
	o.calculateTotal()
	o.displayOrderDetails()
	s.orders = append(s.orders, o)
}

func (s *store) displayAllOrders() {

	fmt.Println("All Orders:")
	for _, order := range s.orders {
		order.displayOrderDetails()
		fmt.Println()
	}
}

func main() {

	store := store{}

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
