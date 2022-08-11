package main

import (
	"fmt"
	"time"
)

// without condition or if
// This construct can be a clean way to write long if-then-else chains.
func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	case t.Hour() > 22:
		fmt.Println("Good Night.")
	default:
		fmt.Println("Good evening.")
	}
}
