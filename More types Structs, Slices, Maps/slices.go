package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4] //low to high-1 (slice and assign to s)
	fmt.Println(s)            //print s

}
