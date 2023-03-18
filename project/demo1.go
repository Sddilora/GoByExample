package project

import (
	"encoding/json"
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
