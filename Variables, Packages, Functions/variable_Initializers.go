package main

import "fmt"

var i, j int = 10, 99

// := does not work explicit

func main() {
	var C, Java, Python = true, false, "Nani!"
	//short assign for implicit for var
	k := 3 // k is an int
	fmt.Println("Output: ", i, j, C, Java, Python)
	fmt.Println("Short assignment variable for k: ", k)
}
