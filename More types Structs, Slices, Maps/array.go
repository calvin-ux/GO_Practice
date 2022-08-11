package main

import "fmt"

func main() {
	//var a [10]int
	//declares a variable a as an array of ten integers.
	var a [2]string // e,g, a[0] and a[1]
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

//An array's length is part of its type, so arrays cannot be resized
