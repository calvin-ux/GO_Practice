package main

import (
	"fmt"
	"os"
	"time"
)

func Select(c chan int, quit chan struct{}) {
	//switch case
	//select for channel async
	for {
		time.Sleep(time.Second * 1)
		select {
		case <-c:
			fmt.Println("RECIVED")
		case <-quit:
			fmt.Println("QUIT")
			os.Exit(1)
		}
	}
}

func main() {
	c := make(chan int) //1 size at a time
	quit := make(chan struct{})

	//unknown
	go func() {
		c <- 1
		quit <- struct{}{}
	}()

	go Select(c, quit) //passing channels too function
	select {}

	fmt.Println("VS CODE")
}
