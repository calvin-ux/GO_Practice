package main

import (
	"fmt"
	"math"
)

// func main() {
// 	var x, y int = 3, 4
// 	var f float64 = math.Sqrt(float64(x*x + y*y))
// 	var z uint = uint(f)
// 	fmt.Println(x, y, z)
// }

// OR using short assignments
func main() {
	a, b := 4, 5
	f := math.Sqrt(float64(a*a + b*b))
	c := uint(f)
	fmt.Println(a, b, c)
}
