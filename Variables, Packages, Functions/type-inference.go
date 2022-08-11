package main

import "fmt"

func main() {
	v := 42           //v is an int!
	f := 3.142        //f is a float64
	g := 0.867 + 0.5i //complex128
	fmt.Printf("v is of typ %T\n", v)
	fmt.Printf("f is of typ %T\n", f)
	fmt.Printf("f is of typ %T\n", g)
}
