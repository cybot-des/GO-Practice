package main

import (
	"fmt"
)

func main() {
	a := func(num1, num2 int) {
		fmt.Println("Adding: ", num1+num2)
	}

	b := func(num1, num2 int) {
		fmt.Println("Subtracting: ", num1-num2)
	}

	demo(a)
	demo(b)

	///////////////
	c := add
	d := subtract
	c(40, 30)
	d(40, 30)
	demo(c)
	demo(d)
}

func demo(fun func(num1, num2 int)) {
	fun(10, 20)
}

func add(num1, num2 int) {
	fmt.Println("Adding: ", num1+num2)
}

func subtract(num1, num2 int) {
	fmt.Println("Subtracting: ", num1-num2)
}
