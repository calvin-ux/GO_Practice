package main

import "fmt"

//https://go.dev/tour/moretypes/8
func main() {
	//normal arrray execution
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)
	//with slice
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)
	//A slice does not store any data, it just describes a section of an underlying array.

	/*
		Changing the elements of a slice
		modifies the corresponding elements of its underlying array.
	*/
	//using slice with pointers
	b[0] = "XVV"
	fmt.Println(a, b)
	fmt.Println(names)
}
