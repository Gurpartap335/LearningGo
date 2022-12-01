package main

import (
	"fmt"
)

func main() {
	arr := []int{5, 4, 3, 2, 1}
	fmt.Println(arr)
	insertion(arr)
	fmt.Println(arr)
	
}

func insertion(arr[] int) {
	for i := 0; i < len(arr) - 1; i++ {
		for j := i + 1; j > 0; j-- {
			if arr[j] < arr[j - 1] {
				arr[j], arr[j - 1] = arr[j - 1], arr[j]
			} else {
				break
			}
	    }
    }
}