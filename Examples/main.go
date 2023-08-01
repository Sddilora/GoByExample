package main

import (
	"fmt"
	"lesson/Examples/project"
	"log"
)

func main() {

	project.GetAllProducts()

	project.AddProduct()
	project.GetAllProducts() // To see if wee succeed

	project.AddProduct2()
	project.GetAllProducts() //To see if wee succeed

	project.AddProduct3()
	project.GetAllProducts() //To see if wee succeed

	products, err := project.GetAllProducts2()
	log.Println(products, err)

	products, _ = project.GetAllProducts2()
	for i := 0; i < len(products); i++ {
		fmt.Println(products[i].ProductName)
	}

	project.AddProduct_2()
	products, _ = project.GetAllProducts2()
	for i := 0; i < len(products); i++ {
		fmt.Println(products[i].ProductName)
	}

	//Output
	//
	// Laptop
	// Mouse
	// Water
	// Cell Phone
	// Microphone
	// Microphone
	// Keyboard
}
