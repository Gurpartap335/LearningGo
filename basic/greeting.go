package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Enter your name")
	var name string
	fmt.Scanln(&name)
	name = strings.TrimSpace(name)
	fmt.Printf("Hello, %s! I am Go\n", name)
}
