package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting...")

	arr := []int{7, 19, 12, 90, 34, 67, 22, 84, 20}
	fmt.Printf("Addition is %d\n", AddElements(arr))

	fmt.Printf("Max element is %d\n", FindMax(arr))
	fmt.Printf("Min element is %d\n", FindMin(arr))

	fmt.Println(searchElement(arr, 67))

	sortArray(arr)

}

func AddElements(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func FindMax(arr []int) int {
	max := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func FindMin(arr []int) int {
	min := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}

func searchElement(arr []int, element int) string {
	flag := false
	var pos int
	for i := 0; i < len(arr); i++ {
		if arr[i] == element {
			flag = true
			pos = i
			break
		}
	}
	if flag == true {
		return fmt.Sprintf("Found at %d", pos)
	}
	return "Not found"
}

func sortArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}

	fmt.Println(arr)
}
