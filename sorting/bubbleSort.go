package main

import (
	"fmt"
)

func main() {
	arr := []int{5, 4, 3, 2, 1,}
	fmt.Println(arr)

	bubble(arr)
	fmt.Println(arr)
}

// bubble sort sort the array by swaping adjacent elements if they are in
// wrong order.  Largest element comes last in the array.
func bubble(arr []int) {
	var swapped bool
	for i := 0; i < len(arr) - 1; i++ {
		swapped = false 
		for j := 0; j < len(arr) - i - 1; j++ {
			if (arr[j] > arr[j + 1]) {
				arr[j + 1], arr[j] = arr[j], arr[j + 1]
				swapped = true
			}
		}

		if (!swapped) {
			break
		}
	}
}