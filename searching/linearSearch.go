package main

import (
	"fmt"
)

func main() {

	arr := []int{1, 2, 3, 4, 11, 3, 3, 13, 5}
	a := []int{}
	fmt.Println(linearAll(arr, 3, 0, a))
}

func searchL(arr []int, target int) int {

	for i := 0; i < len(arr); i++ {
		if (arr[i] == target) {
			return i;
		}
	}
	return -1
}

func linear(array []int, query int) (int) {
	for i, item := range array {
		if item == query {
			return i
		}
	}
	return -1
}

// return first index of target element or -1 if doesnot exit
func linearR(arr []int, target int, index int) int {
	if (index == len(arr)) {
		return -1
	}
	if (arr[index] == target) {
		return index
	} else {
		return linearR(arr, target, index + 1)
	}
}

// return all the index of element
func linearAll(arr[] int, target int, index int, ans []int) []int {
	if (index == len(arr)) {
		return ans
	}

	if (arr[index] == target) {
		ans = append(ans, index)
	} 

    return linearAll(arr, target, index + 1, ans)
}

// return value in boolean if value exits or not.
func linearR2(arr []int, target int, index int) bool {
	if (index == len(arr)) {
		return false
	}

	return arr[index] == target || linearR2(arr, target, index + 1)
}
