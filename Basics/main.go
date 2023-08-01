package main

import (
	"fmt"
	"lesson/Basics/defer_statement"
	"lesson/Basics/error_handling"
	"lesson/Basics/restful"
	"lesson/Basics/string_functions"
)

func main() {

	defer_statement.TestA2()
	defer_statement.Demo3()

	error_handling.Demo1()
	error_handling.Demo2()

	error_handling.Demo3()
	fmt.Println(error_handling.GuessWhat2(102)) //konsol :: 102---Out of range

	string_functions.Demo1()
	string_functions.Demo2()

	restful.Demo1()
	restful.Demo2()

}
