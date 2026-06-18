package main

import (
	"fmt"
)

func main() {
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}

	fmt.Println(arr1)
	fmt.Println(arr2)

	// This example declares two arrays (arr1 and arr2) with inferred lengths
	var arr3 = [...]int{1, 2, 3}
	arr4 := [...]int{4, 5, 6, 7, 8}

	fmt.Println(arr3)
	fmt.Println(arr4)

	// This example initializes only the second and third elements of the array
	arr5 := [5]int{1: 10, 2: 40}
	fmt.Println(arr5)

	arr6 := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	fmt.Println(len(arr6))
}
