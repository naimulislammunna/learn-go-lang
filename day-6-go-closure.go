package main

import "fmt"

const a = 10

var b = 200

func outer() func() {
	money := 100
	age := 30
	fmt.Println("age:", age)

	show := func() {
		money = money + a + b
		fmt.Println("money:", money)
	}
	return show
}

func call() {
	increment1 := outer()
	increment1()
	increment1()
	increment1()

	increment2 := outer()
	increment2()
	increment2()
	increment2()
}

func main() {
	call()
}

func init() {
	fmt.Println("init")
}
