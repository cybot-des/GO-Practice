package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	defer thankyou()
	student := Student{}
	fmt.Println("========= WELCOME =========")
	buff := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Name: ")
	student.Name, _ = buff.ReadString('\n')
	fmt.Println("Enter Class: ")
	data, _, _ := buff.ReadLine()
	student.Class = string(data)
	student.Name = strings.TrimSpace(student.Name)
	fmt.Printf("%+v \n", student)

	if student.Class == "mca" {
		fmt.Println(" Hello World!")
	}
}

type Student struct {
	Name  string
	Class string
}

func thankyou() {
	fmt.Println("======= ThankYou ========")
}
