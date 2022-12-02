package main

import (
	"fmt"
)

// https://stackoverflow.com/questions/35707222/swap-two-numbers-golang

func main() {
	// a, b := 10, 5
	// fmt.Println("a: ", a, "b: ", b)
	// a, b = b, a
	// fmt.Println("a: ", a, "b: ", b)

	fmt.Println(myFunction())

}

func myFunction() []int {
	a, b := 10, 5
	a, b = b, a
	return []int{a, b}
}