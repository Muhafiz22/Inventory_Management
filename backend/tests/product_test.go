package tests_test

import (
	"fmt"
	"inventory_management/backend/models"
	"inventory_management/backend/services"
	"testing"
)

/*TO-DO
1.first make arrangements for calling get function
*/

func TestAddProduct(t *testing.T) {

	ProductService := &services.ProductService{}
	products := []models.Product{} //initialising the slice of struct else it may create a null pointer error

	//productMethods := &models.Product{}
	products = append(products, models.Product{ID: 1, Name: "kitkat", Price: 20.0, Quantity: 10})
	//productMethods.Get(products)
	fmt.Println("\nAbout to add Product")                          //Debug Statement
	updatedProductList, err := ProductService.AddProduct(products) //didn't use the updatedProductList return of AddProduct function.
	if err != nil {
		fmt.Println("Error Adding Product:", err)
	} else {
		fmt.Println(updatedProductList)
	}
}
