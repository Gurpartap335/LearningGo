package main

import "fmt"

const favColor string = "blue"

func main() {
	var guess string
	// Create an input loop
	for {
		// Ask user to enter favorite color
		fmt.Println("Guess my favorite color: ")
		if _, err := fmt.Scanln(&guess); err != nil {
			fmt.Println("%s\n", err)
			return
		}

		if favColor == guess {
			fmt.Printf("%q is my favorite color!\n", favColor)
			return
		}
		fmt.Printf("Wrong")
	}
}
