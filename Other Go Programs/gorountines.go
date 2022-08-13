package main

import (
	"fmt"
	"time"
)

func heavy() {
	//time.Sleep(time.Second * 5) //execute prog and display after 5 secs

	for {
		time.Sleep(time.Second * 1)
		fmt.Println("Heavy")
	}
}

func superHeavy() {
	for {
		time.Sleep(time.Second * 2)
		fmt.Println("Super Heavy")
	}
}

func main() {
	//heavy()
	//NOTE: If parent class complete then rest of the go routines won't execute hence select {} or time.Sleep
	//asyn func
	go heavy() //goroutine is the way we call func and execute in background
	go superHeavy()
	fmt.Println("Vs-code")
	//time.Sleep(time.Second * 5) //run for addtional 5 secods before exit
	//OR
	select {} //run till ctrl+c
}
