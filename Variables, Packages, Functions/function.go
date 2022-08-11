package main

import (
	"fmt"
)

// function with 0 or more args
func add(x int, y int) int {
	return x + y
}

// shortened parameter with same type (int)
func sub(a, b int) int {
	return a - b
}
func mul(a, b, c int) int {
	return a * b * c
}

func main() {
	fmt.Println(add(43, 97))
	fmt.Println(sub(56, 12))
	fmt.Println(mul(12, 9, 2))
}
