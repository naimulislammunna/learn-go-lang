package main

import (
	"fmt"
)

func main() {
	/* var : You always have to specify either type or value
	:= It is not possible to declare a variable using :=, without assigning a value to it. */

	var student1 string = "John" //type is string
	var student2 = "Jane"        //type is inferred
	x := 2                       //type is inferred

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)

	var a string
	var b int
	var c bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
