package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var a [10]int

	fmt.Println(len(a))

	fmt.Println("Enter array element: ")
	for i := 0; i < 10; i++ {
		fmt.Scan(&a[i])
	}

	fmt.Println("Printing array element: ")
	for i := 0; i < 10; i++ {
		fmt.Println(a[i])
	}
}