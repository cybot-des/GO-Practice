package main

import (
	"fmt"
	"time"
)

func main() {
	go printMessage()
	go printMessage()
	go printMessage()
	go printMessage()
	go printMessage()
	go printMessage()
	fmt.Println("Hello, world!")
	Hello()
	time.Sleep(2 * time.Second)
}

func Hello() {
	time.Sleep(2 * time.Second)
	fmt.Println("Hello is called!")
}

func printMessage() {
	for i := 0; i <= 10; i++ {
		fmt.Println("Printed Index: ", i)
		time.Sleep(1 * time.Second)
	}
}
