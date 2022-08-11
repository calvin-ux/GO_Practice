package main

import "fmt"

//https://go.dev/tour/moretypes/4
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v   //point to v
	p.X = 1e9 //Struct fields can be accessed through a struct pointer.
	fmt.Println(v)
}

/*
 we could write (*p).X. However, that notation is cumbersome
*/
