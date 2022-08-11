package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

//can cause harm in  longer function
//used in short function
// "naked" return --> statement without args nor values e.g. return
func main() {
	fmt.Println(split(17))
}
