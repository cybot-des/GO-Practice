package main

import (
	"fmt"
)

func main() {

	//-----POINTER-----
	// x := 100
	// var p *int
	// p = &x
	// fmt.Println("Value stored in x = ", x)
	// fmt.Println("Address of x = ", &x)
	// fmt.Println("Value stored in variable p = ", p)

	//----STRUCTURE------

	stud := student{name: "Pooja", class: "MCA", Age: 22,
		ad: Address{Line1: "NMIMS", City: "Mumbai", State: "Maharashtra", Pincode: 400066}}

	// //fmt.Printf("%+v \n", stud)
	// fmt.Printf("%s in class %s having age %d \n", stud.name, stud.class, stud.Age)
	// //fmt.Printf("%+Address ", stud.ad)
	// fmt.Printf("%+v \n", stud)

	var arr [2]student

	arr[0] = stud
	arr[1] = student{name: "ABC"}

	// fmt.Println(arr[0])
	// fmt.Println(arr[1])

	// for i := 0; i < len(arr); i++ {
	// 	fmt.Printf("%+v ", arr[i])
	// }

	// for v := range arr {
	// 	fmt.Printf("%+v ", v)
	// }

	// for i := range arr {
	// 	fmt.Println(i)
	// }

	// for _, v := range arr {
	// 	fmt.Println(v)
	// }

	// var slice []student

	// slice = append(slice, stud)
	// slice = append(slice, student{name: "PQR"})
	// slice = append(slice, student{name: "MNO"}, student{name: "XYZ"})

	// //fmt.Println(slice)
	// index := 4

	// num := []int{1, 2, 3, 4, 5}
	// num = append(num, 6, 7)
	// fmt.Println(num[:])
	// fmt.Println(num[2:6])
	// fmt.Println(num[:3])
	// num = append(num[:index], num[index+1:]...)
	// fmt.Println(num)

	var mapData = make(map[string]string)
	mapData["Name"] = "Pooja"
	mapData["Class"] = "MCA"
	mapData["Year"] = "2023"

	// fmt.Println("All Map Data: ", mapData)
	// fmt.Println()
	// fmt.Println("Name:", mapData["Name"])

	mapAddress := make(map[string]Address)
	mapAddress["Line1"] = Address{Line1: "Andheri"}
	mapAddress["Line2"] = Address{Line1: "Pune"}
	mapAddress["Line3"] = Address{Line1: "Borivali"}
	mapAddress["Line4"] = Address{Line1: "Nashik"}

	// fmt.Println("Map Address: ", mapAddress)

	// _, ok := mapAddress["Line5"]
	// if ok {
	// 	fmt.Println("Valid Address Line")
	// } else {
	// 	fmt.Println("Invalid Address Line")
	// }

	// delete(mapAddress, "Line4")
	// fmt.Println("Map Address: ", mapAddress)

	//------CHECK DUPLICATES------
	// arr1 := []string{"ABC", "PQR", "ABC", "ABC", "MNO", "PQR", "PQR"}
	// mapData1 := make(map[string]int)
	// for _, v := range arr1 {
	// 	value, ok := mapData1[v]
	// 	if ok {
	// 		mapData1[v] = value + 1

	// 	} else {
	// 		mapData1[v] = 1

	// 	}
	// }

	// for i, v := range mapData1 {
	// 	if v > 1 {
	// 		fmt.Printf("%s repeated %d times \n", i, v)
	// 	}
	// }

	// Test()
	// Test1(2)
	// fmt.Println(Test3(1, 2))
	// fmt.Println(Swap(10, 20))
	// fmt.Println(Calculate(10, 20))

	fmt.Println(CalculateVariadic("Pooja", true, 10, 20, 30, 40, 50))
}

type student struct {
	name  string
	class string
	Age   int
	ad    Address
}

type Address struct {
	Line1   string
	City    string
	State   string
	Pincode int
}

//------FUNCTIONS-------

//no arguments no return value
func Test() {
	fmt.Println("Test Function")
}

//with argument no return
func Test1(arg1 int) {
	fmt.Println("Returns ", arg1)
}

//with argument with return
func Test3(arg1 int, arg2 int) int {
	return arg1 + arg2
}

// Multi Return
func Swap(arg1, arg2 int) (int, int) {
	return arg2, arg1
}

//Multi Return Type 2
func Calculate(x, y int) (out1, out2 int) {
	out1 = x + 10
	out2 = y + 20
	return
}

//Variadc Function Return
func CalculateVariadic(name string, status bool, nums ...int) (out int) {
	fmt.Println(" ", nums)
	for i := 0; i < len(nums); i++ {
		fmt.Println("Inputs: ", nums[i])
		out += nums[i]
	}
	fmt.Println("Outputs: ")
	return
}
