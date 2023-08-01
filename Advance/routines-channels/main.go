package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CalculateValue(values chan int) {
	fmt.Println("Calculating Value...")
	value := rand.Intn(10)
	fmt.Println("Calculated Random Value: {}", value)
	values <- value
}

func CalculateValue2(c chan int) {
	fmt.Println("Calculating Value 2...")
	value := rand.Intn(10)
	fmt.Println("Calculated Random Value 2: {}", value)
	time.Sleep(1000 * time.Millisecond)
	c <- value
	fmt.Println("Only Executes after another goroutine performs a receive on the channel")
}

func main() {
	fmt.Println("Go Channel Tutorial")

	//Note - We used make when instantiating our values channel as, like maps and slices, channels must be created before use.
	values := make(chan int) // This call effectively created our new channel so that we could subsequently use it within our CalculateValue goroutine.
	defer close(values)

	go CalculateValue(values)

	value := <-values
	fmt.Println("value is equal to", value)

	valueChannel := make(chan int)
	defer close(valueChannel)

	go CalculateValue2(valueChannel)
	go CalculateValue2(valueChannel)

	values2 := <-valueChannel
	fmt.Println("value 2 is equal to", values2)

}
