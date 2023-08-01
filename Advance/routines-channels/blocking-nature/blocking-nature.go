package main

import (
	"fmt"
	"time"
)

func senderRepeat(c chan string) {
	t := time.NewTicker(time.Second * 1)
	for {
		<-t.C
	}
}
func main() {
	stopChannel := make(chan string)
	c := make(chan string)
	go senderRepeat(c)
	go func() {
		time.Sleep(time.Second * 5)
		fmt.Println("Running...")
		stopChannel <- "stop"
		fmt.Println("stop channel writes are non blocking")
	}()
	fmt.Println(<-c)
}
