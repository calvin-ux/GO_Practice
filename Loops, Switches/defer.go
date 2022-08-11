package main

import "fmt"

//https://go.dev/tour/flowcontrol/12

func main() {
	defer fmt.Println("Hello")

	fmt.Println("World")
}
