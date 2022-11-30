package main

import (
	"fmt"
)

func main() {
	printTraingle(4)
	printTriangleN2(5)
}

func printTraingle(n int) {

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

// with numbers
// 1 2 3 4
// 1 2 3
// 1 2
// 1

func printTriangleN2(n int) {

	for i := 0; i < n; i++ {
		for j := 1; j <= n - i; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}