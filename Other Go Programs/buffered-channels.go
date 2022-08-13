package main

import "fmt"

type Car struct {
	Model string
}

func main() {
	c := make(chan *Car, 3)

	go func() {
		c <- &Car{"1"}
		c <- &Car{"2"}
		c <- &Car{"3"}
		c <- &Car{"4"}
		// c <- 1
		// c <- 2
		// c <- 3
		// c <- 4
		close(c) //close channel at 4
	}()

	for i := range c {
		fmt.Println(i.Model) //print every value
	}
}
