package main

import (
	"fmt"
	"math/rand"
	"time"

)
func main() {

	secret := getRandomNumber()

	

	for matching := false; !matching; {
		guess := getUserInput()

	    matching = isMatching(secret, guess)
	}
}

func getRandomNumber() int {
	rand.Seed(time.Now().Unix())
	return rand.Int() % 11
}

func getUserInput() int {
	fmt.Println("Please input your number")
	var input int
	
	// handle error
	_, err := fmt.Scan(&input) // err var data type is error
	if err != nil {
		fmt.Println("Failed to parse your input")
	} else {
		fmt.Println("You guess: ", input)
	}
	return input
}

func isMatching(secret int, guess int) bool {
	if secret > guess {
		fmt.Println("You guess is too smaller")
		return false
	} else if guess > secret {
		fmt.Println("Your guess is too big")
		return false
	} else {
		fmt.Println("Guess is right.", "\nSecret number: ", secret, "Input number: ", guess)
		return true
	}
}
