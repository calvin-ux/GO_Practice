package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// method with reciver, typr Vertex for v
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
} //2500 = 50

// Methods with pointer receivers can modify the value to which the receiver points (as Scale does here).
func (v *Vertex) Scale(f float64) { //f is 10
	v.X = v.X * f
	v.Y = v.Y * f
} // do scale on 3, 4 to 30, 40

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs()) // 50 now output
}
