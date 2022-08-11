package main

import "fmt"

type Vertex struct {
	X, Y int
}

//https://go.dev/tour/moretypes/5
var (
	v1 = Vertex{1, 2} // has type Vertex
	//v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, p, v3)
}
