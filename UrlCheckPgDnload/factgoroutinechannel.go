package main

import "fmt"

func factorial(n int, ch chan int) {
	fact := 1
	for i:= 2; i <= n; i++ {
		fact *= i
	}
	ch <- fact
}

func main() {
	ch := make(chan int)
	defer close(ch)

	// Call a normal goroutine
	go factorial(5, ch)
	fact := <- ch
	fmt.Printf("The factorial of %d is %d\n", 5, fact)

	// Calling an anonymous goroutine
	go func(n int, c chan int) {
		fact := 1
		for i:= 2; i <= n; i++ {
			fact *= i
		}
		c <- fact
	}(5, ch)     // This is the values to pass to the call
	fact = <- ch
	fmt.Printf("The factorial of %d is %d\n", 5, fact)
}
