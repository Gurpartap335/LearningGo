package main

import (
	"fmt"
)

func main() {
	arr := []int{3, 4, 1, 9, 22, 11, 92}
	fmt.Println(arr)
	selection(arr)
	fmt.Println(arr)
}

func selection(arr []int) {
	for i := 0; i < len(arr) - 1; i++ {
		iMin := i;
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[iMin] {
				iMin = j;
			}
		}

		if iMin != i {
			arr[i], arr[iMin] = arr[iMin], arr[i]
		}
	}
}