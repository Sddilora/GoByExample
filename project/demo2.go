package project

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func GetAllProducts2() ([]Product, error) {

	response, err := http.Get("http://localhost:3000/products")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var products []Product
	json.Unmarshal(bodyBytes, &products) //if we don't put "&" (memory address) before result storage it will not assign the unmarshalled datas to variable "products []Product".
	return products, nil
}

//Output
//
// 2023/03/18 14:11:46 [{1 Laptop 1 5000.99} {2 Mouse 1 50.99} {3 Water 2 3.99} {4 Cell Phone 1 10000} {5 Microphone 1 2000} {6 Microphone 1 2000}] <nil>

func AddProduct_2() (Product, error) {
	product := Product{ProductName: "Keyboard", CategoryId: 1, UnitPrice: 300.0}
	jsonProduct, err := json.Marshal(product)
	if err != nil {
		return Product{}, err
	}
	response, err := http.Post("http://localhost:3000/products", "application/json;charset=utf-8", bytes.NewBuffer(jsonProduct))
	if err != nil {
		return Product{}, err
	}

	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}

	var productResponse Product

	json.Unmarshal(bodyBytes, &productResponse)

	return productResponse, nil

}
