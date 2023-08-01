package main

import "fmt"

type MyFunction func(int, int) int

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func doOperation(operation MyFunction, a, b int) int {
	return operation(a, b)
}

////////////////////////////////////////////////////////////

func higherOrder(fn func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		return fn(a, b) + 100
	}
}

func add2(a, b int) int {
	return a + b
}

func main() {
	result1 := doOperation(add, 10, 5)
	result2 := doOperation(multiply, 10, 5)

	fmt.Println("Addition Result:", result1)       // Output: Addition Result: 15
	fmt.Println("Multiplication Result:", result2) // Output: Multiplication Result: 50

	addWithBonus := higherOrder(add2)
	result := addWithBonus(10, 5)

	fmt.Println("Result with Bonus:", result) // Output: Result with Bonus: 115
}
