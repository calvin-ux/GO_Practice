package main

import (
	"fmt"
	"sync"
)

func main() {

	//wait group--->property of sync
	wq := &sync.WaitGroup{}
	//add 2 func/routines, done, wait
	wq.Add(2)

	// mutexes(critical section) only for 1 go to run for accessing shared resoureces

	//aynononmous or lamdba function
	go func() {
		fmt.Println("Hello") //1st
		wq.Done()
	}() //call function
	go func() {
		fmt.Println("World") //2n
		wq.Done()
	}()

	fmt.Println("Fin")
	wq.Wait()
	fmt.Println("EXIT")
}
