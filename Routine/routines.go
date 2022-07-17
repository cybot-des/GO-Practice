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
	// wg.Add(1)
	// go printMsg()
	// fmt.Println("===STARTED WAITING===")
	// wg.Wait()
	// fmt.Println("===ENDED WAITING===")

	//-------EXAMPLE 2 OF ROUTINES USING WAIT GROUP
	// wg.Add(2)
	// go printMsg()
	// go printMsg()
	// fmt.Println("===STARTED WAITING===")
	// wg.Wait()
	// fmt.Println("===ENDED WAITING===")

	//-------EXAMPLE 3 OF ROUTINES USING WAIT GROUP
	// wg.Add(3)
	// go printMsg()
	// go printMsg()
	// fmt.Println("===STARTED WAITING===")
	// wg.Wait() //Endlessly stuck in waiting loop( deadlock )
	// fmt.Println("===ENDED WAITING===")

	// ROUTINES USING CHANNELS

	//------EXAMPLE 1 OF ROUTINES USING CHANNELS
	// ch := make(chan string)
	// go printMsgch(ch)

	// fmt.Println("Processing....")
	// data := <-ch
	// fmt.Println("Recieved Data: ", data)

	//-----EXAMPLE 2 OF ROUTINES USING CHANNELS (MULTI CHANNELS)
	ch3 := make(chan int)
	ch4 := make(chan int)

	go DemoSequence(ch3)
	go DemoSequence(ch4)

	for {
		select {
		case data := <-ch3:
			{
				fmt.Println("Data from Channel 3: ", data)
			}

		case data := <-ch4:
			{
				fmt.Println("Data from Channel 4: ", data)
			}
		}

	}

}

func DemoSequence(ch chan int) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		ch <- i
	}
	ch <- 100
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

func printMsgch(ch chan string) {
	time.Sleep(2000 * time.Millisecond)
	ch <- "Hello World"
}
