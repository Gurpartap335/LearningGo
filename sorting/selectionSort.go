package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	selectionD(arr)
	fmt.Println(arr)
}

// ascending order
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

// descending order
func selectionD(arr []int) {
	for i := 0; i < len(arr) - 1; i++ {
		iMax := i;
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > arr[iMax] {
				iMax = j;
			}
		}

		if iMax != i {
			arr[i], arr[iMax] = arr[iMax], arr[i]
		}
	}
}