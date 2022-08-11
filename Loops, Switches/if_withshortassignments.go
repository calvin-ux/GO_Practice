package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim { //if cond + and
		return v
	}
	return lim
}
func main() {
	fmt.Println(
		pow(7, 2, 50),  //for 9 --(3, 2, 10)
		pow(7, 3, 200), //for 27 -- (3, 3, 20)
	)
}
