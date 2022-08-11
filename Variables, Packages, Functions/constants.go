package main

import "fmt"

const Pi = 3.16

func main() {

	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")
	//Constants cannot be declared using the := syntax
	const Truth = true
	fmt.Println("Go Rules ?", Truth)
}
