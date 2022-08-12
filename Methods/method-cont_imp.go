package main

import (
	"fmt"
	"math"
)

// You can declare a method on non-struct types, too.
type MyFloat float64

// You can only declare a method with a receiver whose type is defined in the same package as the method
// In this example we see a numeric type MyFloat with an Abs method.
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
