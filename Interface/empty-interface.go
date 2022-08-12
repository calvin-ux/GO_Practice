// https://go.dev/tour/methods/14
// Empty interfaces are used by code that handles values of unknown type
package main

import "fmt"

func main() {
	var i interface{}
	describe(i)

	i = 365
	describe(i)

	i = "Hola"
	describe(i)
}

//The interface type that specifies zero methods is known as the empty interface--->interface{}
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
