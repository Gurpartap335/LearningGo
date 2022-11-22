package main

import "fmt"

func main() {

	var size int
	fmt.Println("Enter the size of array")
	fmt.Scan(&size)
	var arr [size]int

	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}

	for i := 0; i < size; i++ {
		fmt.Println(arr[i])
	}

}