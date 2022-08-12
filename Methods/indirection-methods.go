package main

import "fmt"

type Vertex struct {
	X, Y float64
}

//method with reciver(v) and pointer
func (v *Vertex) Scale(f float64) { //f is 2 for scale
	v.X = v.X * f
	v.Y = v.Y * f
}

//function/method with pointer
func ScaleFunc(v *Vertex, f float64) { //scales by *10
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {

	//pass and scale normal auto call from pointer reciver
	v := Vertex{3, 4}
	v.Scale(2) //Go interprets the statement v.Scale(5) as (&v).Scale(5)
	//since the Scale method has a pointer receiver.
	ScaleFunc(&v, 10)

	//pointing to normally
	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
