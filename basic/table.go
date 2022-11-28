package main

import (
	"fmt"
)

// Program for printing a number table
func main() {
	var n int
	fmt.Println("Enter a number: ")
	fmt.Scan(&n)

	for i := 1; i < 11; i++ {
		fmt.Print(n * i , " ")
	}
}