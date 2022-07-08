package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting...")
	fmt.Println(Add(3, 9))
	fmt.Println(Sub(4, 7))
	fmt.Println(Mul(8, 2))
	fmt.Println(Div(5, 13))
}

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}

func Div(a, b float32) float32 {
	return a / b
}
