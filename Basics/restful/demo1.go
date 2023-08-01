package restful

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"user_id"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// Http GET Request
func Demo1() {
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Print(errors.New("there is an error while generating http response"))
	}
	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Print(err)
	}

	bodyString := string(bodyBytes)

	var todo Todo
	json.Unmarshal(bodyBytes, &todo)
	fmt.Println(todo)
	fmt.Println(bodyString)
}

//  Output:
//
//	->{0 1 delectus aut autem false}
// 	  {
// 		"userId": 1,
// 		"id": 1,
// 		"title": "delectus aut autem",
// 		"completed": false
//    }

// Http Post Request
func Demo2() {
	todo := Todo{1, 2, "Shopping will be done", false}
	//turn the value to the json format
	jsonTodo, err := json.Marshal(todo)
	if err != nil {
		log.Print(err)
	}

	response, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json;charset=utf-8", bytes.NewBuffer(jsonTodo))
	if err != nil {
		log.Print(err)
	}
	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Print(err)
	}

	bodyString := string(bodyBytes)

	var todoResponse Todo
	json.Unmarshal(bodyBytes, &todoResponse)
	fmt.Println(todoResponse)
	fmt.Println(bodyString)

}

//Output:
//
// {1 201 Shopping will be done false}
// {
//   "user_id": 1,
//   "id": 201,
//   "title": "Shopping will be done",
//   "completed": false
// }
