package main

import "fmt"

func main() {

	//---------------VARIABLES-------------------
	// var name string = "Pooja"
	// var age int = 22
	// var email string = "user@gmail.com"
	// city, state := "Mumbai", "Maharashtra"
	// var isEligible bool = true

	// fmt.Println("Name: ", name, " Age: ", age, " Email: ", email)
	// fmt.Println("State: ", state, "City: ", city, " Eligiblity: ", isEligible)
	// reflect.TypeOf(isEligible)

	// name, surname := os.Args[1], os.Args[2]
	// fmt.Println("Name: ", name, "  Surname: ", surname)

	// io := bufio.NewReader(os.Stdin)
	// fmt.Println("Enter Name: ")
	// name, err := io.ReadString('\n')

	// if err != nil {
	// 	fmt.Println("Error: ", err)

	// }

	// fmt.Println("Enter Age: ")
	// age, err := io.ReadString('\n')
	// if err != nil {
	// 	fmt.Println("Error: ", err)

	// }
	// fmt.Println("You Entered:", name, "\t", age)

	//-----CONDITION LOOP FOR --------
	// var sum = 0
	// for i := 0; i <= 10; i++ {
	// 	sum++
	// }
	// fmt.Println("Sum: ", sum)

	// // for {
	// // 	fmt.Printf("Infinite loop")
	// // }

	// for j := 0; j < 10; j++ {
	// 	if j == 4 {
	// 		continue
	// 	}
	// 	if j == 8 {
	// 		break
	// 	}
	// 	fmt.Println("Index: ", j)
	// }

	//-------IF ELSE STATEMENT---------

	// io := bufio.NewReader(os.Stdin)
	// day, err := io.ReadString('\n')
	// if err != nil {
	// 	fmt.Println("Error reading")
	// }
	// //fmt.Println("Day", day, "Hello")

	// day = day[:len(day)-1]

	// if day == "1" {
	// 	fmt.Println("Monday")
	// } else if day == "2" {
	// 	fmt.Println("Tuesday")
	// } else if day == "3" {
	// 	fmt.Println("Wednesday")
	// } else {
	// 	fmt.Println("Random Day")
	// }

	// if i := (20%2 == 0); i {
	// 	fmt.Println("Even Number")
	// } else {
	// 	fmt.Println("Odd Number")
	// }

	// switch os := runtime.GOOS; os {
	// case "windows":
	// 	fmt.Println("Running on Windows")

	// case "Linux":
	// 	fmt.Println("Running on Linux")

	// case "Mac OS X":
	// 	fmt.Println("Running on Mac OS X")

	// default:
	// 	fmt.Println("Running on unknown OS")
	// }

	// for day := 0; day <= 10; day++ {
	// 	switch day {
	// 	case 0, 2, 4, 6, 8:
	// 		fmt.Println("Even Day")
	// 	case 1, 3, 5, 7, 9:
	// 		fmt.Println("Odd Day")
	// 	default:
	// 		fmt.Println("Unknown Day")
	// 	}
	// }

	//-------Defer Keyword-------

	defer PrintIndex(22)

	for i := 0; i < 10; i++ {
		defer PrintIndex(i)
		//fmt.Println("Position: ",i);
	}

	defer PrintIndex(89)

}

func PrintIndex(index int) {
	fmt.Println("Index: ", index)
}
