package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("when is Friday?")
	today := time.Now().Weekday()
	switch time.Friday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In 2 days")
	case today + 3:
		fmt.Println("In 3 days")
	default:
		fmt.Println("Too far away")
	}
}
