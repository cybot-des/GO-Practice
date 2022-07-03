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

	fmt.Println("Map Address: ", mapAddress)

	// _, ok := mapAddress["Line5"]
	// if ok {
	// 	fmt.Println("Valid Address Line")
	// } else {
	// 	fmt.Println("Invalid Address Line")
	// }

	delete(mapAddress, "Line4")
	fmt.Println("Map Address: ", mapAddress)

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
