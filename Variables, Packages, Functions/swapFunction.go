package main

import "fmt"

func swap(x, y, z string) (string, string, string) {
	return z, y, x
}

func main() {
	a, b, c := swap("General Kenobi", "There", "Hello")
	fmt.Println(a, b, c)
}
