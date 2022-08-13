package main

import (
	"fmt"
	"time"
)

func main() {
	//var c chan int //decalaration
	c := make(chan int)
	//<var name>, chan keyword, and <datatype>

	//send in goroutine
	go func() {
		//send based on same <datatype> of channel c
		c <- 1
	}()

	//sniff
	val := <-c //recieve
	fmt.Println(val)

	//send in goroutine(SAME THING ABOVE REPILICATED)
	go func() {
		c <- 2
	}()
	//sniff
	time.Sleep(time.Second * 2)
	val = <-c //alredy done above use only =
	fmt.Println(val)

	fmt.Println(c)
}
