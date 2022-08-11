package main

import "fmt"

//The zero value of a slice is nil.
func main() {
	var s []int
	fmt.Println(s, len(s), cap(s)) //s[] lenght and capacity
	if s == nil {
		fmt.Println("nil!")
	}
}

//A nil slice has a length and capacity of 0 and has no underlying array.
