package main

import (
	"fmt"
)
/*
*
*	*
*	* 	*
*	*	*	*
*/
func main() {
	printTriangle(4)
	printTriangleN1(5)
}

func printTriangle(n int) {

	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	} 
}

// with Numbers
func printTriangleN1(n int) {

	for i := 1; i <= n; i++ {
		for j := 1; j <= i + 1; j++ { 
			fmt.Print(j, " ")
		}
		fmt.Println()
	} 
}