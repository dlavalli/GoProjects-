package tools

import "fmt"

// Iterates through the characters of a string
func PrintStringChars(s string) {
	r :=[]rune(s)
	for _, ch := range r {
		fmt.Printf("A character: %c (%d)\n", ch, ch);
	}
}

// Iterates through the characters of a string
// and reverse the string
func ReverseString(s string) string {
	sr := []rune(s)  // Convert a string to a rune slice
	r := []rune{}

	for i:=len(s)-1;i >= 0;i-- {
		r = append(r, sr[i])
	}
	return string(r)
}

// In mathematical terms, the sequence Fn of Fibonacci numbers is defined by the recurrence relation
// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, ……..
func FibonacciRecur(index int) int {
	if index <= 1 {
		return index
	}
	return FibonacciRecur(index-1) + FibonacciRecur(index-2)
}

func SumArray(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}
