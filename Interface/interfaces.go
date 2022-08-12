package main

import (
	"fmt"
	"math"
)

type Abuser interface {
	Abs() float64
}

func main() {
	var a Abuser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abuser
	a = &v // a *Vertex implements Abuser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	a = &v

	fmt.Println(a.Abs())
}

// perviuosly done code
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
