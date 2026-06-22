package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// The continue statement is used to skip one or more iterations in the loop
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}

	// The break statement is used to break/terminate the loop execution.

	for p := 0; p < 5; p++ {
		if p == 3 {
			break
		}
		fmt.Println(p)
	}

}
