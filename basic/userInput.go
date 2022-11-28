package main

import (
	"fmt"
)

func main() {
	var firstName string
	var lastName string
	var color string
	var character string

	fmt.Println("Enter Your first name")
	fmt.Scanln(&firstName)

	fmt.Println("Enter Your last name")
	fmt.Scanln(&lastName)

	fmt.Println("What's your favorite color")
	fmt.Scanln(&color)

	fmt.Println("Enter your favorite character name")
	fmt.Scanln(&character)

	fmt.Printf("My name is %s %s. My favorite color is %s and character name is %s.", firstName, lastName, color, character)
}
