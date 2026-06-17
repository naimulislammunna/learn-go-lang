package main

import "fmt"

/* Naming Rules for Go Functions

- A function name must start with a letter
- A function name can only contain alpha-numeric characters and underscores (`A-z`, `0-9`, and `_` )
- Function names are case-sensitive

Parameters and their types are specified after the function name, inside the parentheses

If we (for some reason) do not want to use some of the returned values, we can add an underscore (`_`), to omit this value.

*/

func myFunction(x int, y int) (result int) {
	result = x + y
	return
}

// Go Recursion
func testcount(x int) int {
	if x == 11 {
		return 0
	}
	fmt.Println(x)
	return testcount(x + 1)
}

func main() {
	total := myFunction(1, 2)
	fmt.Println(total)
	testcount(4)
}
