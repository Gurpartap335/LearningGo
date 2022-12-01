package main

import (
	"fmt"
)

func main() {

	arr := []int{1, 2, 3, 4, 5, 23, 122}

	fmt.Print(searchB(arr, 23))
	fmt.Println(binaryR(arr, 122, 0, len(arr) - 1))
}

func searchB(arr []int, target int) int {

	s := 0
	e := len(arr) - 1

	for ; s <= e; {
		m := s + (e - s) / 2
		if arr[m] == target {
			return m
		} else if arr[m] < target {
			s = m + 1
		} else {
			e = m - 1
		}
	}

	return -1
}

func binaryR(arr []int, target int, s int, e int) (int) {
	if (s > e || len(arr) == 0) {
		return -1
	}

	m := s + (e - s) / 2
	if (arr[m] == target) {
		return m
	} else if (arr[m] < target) {
		return binaryR(arr, target, m + 1, e)
	} else {
		return binaryR(arr, target, s, m - 1)
	}
}