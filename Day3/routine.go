package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	fmt.Println("Enter Age: ")
	age, _ := buff.ReadString('\n')
	age1, err := strconv.Atoi(string(strings.TrimSpace(age)))
	if err != nil {
		age1 = 0
	}
	student.age = age1
	fmt.Println("Is age valid? ", student.isValid())
	student.printStudent()
}

type Student struct {
	Name  string
	Class string
	age   int
}

type Teacher struct {
}

//METHODS

func (s Student) printStudent() {
	fmt.Println("Student ", s.Name, " class ", s.Class, " Age ", s.age)
}

func (s Student) isValid() bool {
	return s.age >= 18
}

//Overloading like by replacing reciever
func (t Teacher) isValid() bool {
	return true
}

func thankyou() {
	fmt.Println("======= ThankYou ========")
}
