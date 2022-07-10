package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	//-------EXAMPLE OF ROUTINES USING TIME.SLEEP
	// go printMessage()
	// go printMessage()
	// go printMessage()
	// go printMessage()
	// go printMessage()
	// go printMessage()
	// fmt.Println("Hello, world!")
	// Hello()
	// time.Sleep(2 * time.Second) ---TEMPORARY SOLUTION

	//-------EXAMPLE 1 OF ROUTINES USING WAIT GROUP
	wg.Add(1)
	go printMsg()
	fmt.Println("===STARTED WAITING===")
	wg.Wait()
	fmt.Println("===ENDED WAITING===")

	//-------EXAMPLE 2 OF ROUTINES USING WAIT GROUP
	wg.Add(2)
	go printMsg()
	go printMsg()
	fmt.Println("===STARTED WAITING===")
	wg.Wait()
	fmt.Println("===ENDED WAITING===")

	//-------EXAMPLE 3 OF ROUTINES USING WAIT GROUP
	wg.Add(3)
	go printMsg()
	go printMsg()
	fmt.Println("===STARTED WAITING===")
	wg.Wait() //Endlessly stuck in waiting loop( deadlock )
	fmt.Println("===ENDED WAITING===")

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

func printMsg() {
	defer wg.Done()
	for i := 0; i <= 10; i++ {
		fmt.Println("Printed Index: ", i)
		time.Sleep(1 * time.Second)
	}
}
