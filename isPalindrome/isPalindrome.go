package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	num := x
	result := 0

	for num != 0 {
		r := num % 10
		num = num / 10
		result = result*10 + r
	}

	if result == x {
		return true
	}
	return false
}

func main() {
	x := 12321
	fmt.Println(isPalindrome(x))
}
