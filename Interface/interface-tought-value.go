package main

import (
	"fmt"
	"math"
)

// interface
type I interface {
	M()
}

// struct
type T struct {
	S string
}

// method with pts and reciver
func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

// function
func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
