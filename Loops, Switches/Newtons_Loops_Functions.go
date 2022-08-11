package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	fmt.Printf("Sqrt approximation of %v:\n", x) //x = 2
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)                       // Newton's method + division by it's derivate 2z for z*z
		fmt.Printf("Iteration %v, value = %v\n", i, z) //iter i and val z
	}
	return z

}

func main() {
	fmt.Println(Sqrt(2))
}
