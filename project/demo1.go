package project

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Product struct {
	Id          int     `json:"id"`
	ProductName string  `json:"product_name"`
	CategoryId  int     `json:"category_id"`
	UnitPrice   float64 `json:"unit_price"`
}

type Category struct {
	Id           int    `json:"id"`
	CategoryName string `json:"category_name"`
}

func GetAllProducts() {
	response, err := http.Get("http://localhost:3000/products")
	if err != nil {
		log.Print(err)
	}
	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Print(err)
	}

	var products []Product
	json.Unmarshal(bodyBytes, &products) //if we don't put "&" (memory address) before result storage it will not assign the unmarshalled datas to variable "products []Product".
	log.Println(products)
}

//Output
//
//[{1 Laptop 1 5000.99} {2 Mouse 1 50.99} {3 Water 2 3.99}]

func AddProduct() {
	product := Product{Id: 4, ProductName: "Cell Phone", CategoryId: 1, UnitPrice: 10000.0}
	jsonProduct, err := json.Marshal(product)
	if err != nil {
		log.Print(err)
	}
	response, err := http.Post("http://localhost:3000/products", "application/json;charset=utf-8", bytes.NewBuffer(jsonProduct))
	if err != nil {
		log.Print(err)
	}

	defer response.Body.Close()

	fmt.Println("Product Added")
}

//Output
// Product Added
// 2023/03/18 12:43:08 [{1 Laptop 1 5000.99} {2 Mouse 1 50.99} {3 Water 2 3.99} {4 Cell Phone 1 10000}]

func AddProduct2() {
	product := Product{ProductName: "Microphone", CategoryId: 1, UnitPrice: 2000.0}
	jsonProduct, err := json.Marshal(product)
	if err != nil {
		log.Print(err)
	}
	response, err := http.Post("http://localhost:3000/products", "application/json;charset=utf-8", bytes.NewBuffer(jsonProduct))
	if err != nil {
		log.Print(err)
	}

	defer response.Body.Close()

	fmt.Println("Product Added")
}

//Output
//Product Added
// 2023/03/18 12:45:03 [{1 Laptop 1 5000.99} {2 Mouse 1 50.99} {3 Water 2 3.99} {4 Cell Phone 1 10000} {5 Microphone 1 2000}]

func AddProduct3() {
	product := Product{Id: 4, ProductName: "Microphone 3", CategoryId: 1, UnitPrice: 2000.0}
	jsonProduct, err := json.Marshal(product)
	if err != nil {
		log.Print(err)
	}
	response, err := http.Post("http://localhost:3000/products", "application/json;charset=utf-8", bytes.NewBuffer(jsonProduct))
	if err != nil {
		log.Print(err)
	}

	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}

	var productResponse Product

	json.Unmarshal(bodyBytes, &productResponse)

	fmt.Println("Product Added", productResponse)
}

//Output
//
// Product Added {0  0 0}
// 2023/03/18 12:59:14 [{1 Laptop 1 5000.99} {2 Mouse 1 50.99} {3 Water 2 3.99} {4 Cell Phone 1 10000} {5 Microphone 1 2000} {6 Microphone 1 2000}]
//
//In this output product couldn't be added because of the Id present before. This all about API's structure we can change that while coding our own API.
