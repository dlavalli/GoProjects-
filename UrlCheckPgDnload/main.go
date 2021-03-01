package main

import (
	"fmt"
	f "fmt" // Defines an alias so can import the same file multiple times using an alias
	"mypackages/numbers"
)

func init() {
	fmt.Println("Initializing package main")
}

func main() {
	var x uint = 13

	// Call my sample package function
	even := numbers.Even(x)
	fmt.Printf("%d is even: %v", x, even)

	// Package Define Constants are only accessible if they start with an upper case letter
	f.Printf("The boiling point temperature is: %d", numbers.BoilingPoint)
}
