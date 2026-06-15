package main

import (
	"fmt"
)

func main() {
	// single line comment
	/* this is a comment
	   multy line comment */
	fmt.Println("This is munna")

	var name string = "naimul islam munna"
	age := 25
	var doWork bool
	a := 8
	doWork = false
	fmt.Println(name, age, doWork, a)

	// multiple variable declaration
	var g, h, i = 4, false, "dhaka"
	fmt.Println(g, h, i)

	// block variable
	var (
		d string = "gulsan"
		e int
		f bool = true
	)
	fmt.Println(d, e, f)

	const k int = 88
	const b = 99
	fmt.Println(a, b)

	// The brackets in the else statement should be like } else {:
	age = 39

	if age < 18 {
		fmt.Println("they are child")
	} else if age == 18 || age <= 30 {
		fmt.Println("they are adult")
	} else if age > 30 {
		fmt.Println("they are men or women")
	}

	// switch
	// 	All the case values should have the same type as the switch expression.
	// Otherwise, the compiler will raise an error:
	// It is possible to have multiple values for each case in the switch statement:

	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Not a weekday")
	}

}
