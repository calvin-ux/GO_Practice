package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         //point to i = 42
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         //point to j = 2701
	*p = *p / 37   //divide j
	fmt.Println(j) //see new value of j
}
