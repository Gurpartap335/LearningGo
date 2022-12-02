package main

import "fmt"

func main() {

	// array vs slice
	// arr1 := [3]int {3, 9, 1} // array
	// fmt.Println(arr1)

	// arr2 := []int {1, 2, 3, 4, 5} // slice
	// fmt.Println(arr2)
	// fmt.Println(len(arr2))

	// cities := []string {"London", "NYC", "Colombo", "Tokyo"}
	// fmt.Println(cities)

	// addCity := append(cities, "Auckland")
	// fmt.Println(addCity)
	// fmt.Println(cities)

	num := []int {1, 2, 3, 4, 5}
	fmt.Println(num)

	num = append(num, 12)
	fmt.Println(num)

	num = append(num, 12, 13, 14)
	fmt.Println(num)


	// var a [10]int
	// fmt.Println(len(a))

	// fmt.Println("Enter array element: ")
	// for i := 0; i < 10; i++ {
	// 	fmt.Scan(&a[i])
	// }

	// fmt.Println("Printing array element: ")
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(a[i])
	// }
}