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

}
