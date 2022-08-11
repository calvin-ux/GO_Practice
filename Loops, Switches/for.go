package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

//The init and post statements are optional.
//package main
//import "fmt"
//func main() {
//	sum := 1
//	for ; sum < 1000; {
//		sum += 1
//	}
//	fmt.Println(sum)
//}

//For is Go's "while"
//At that point you can drop the semicolons: C's while is spelled for in Go.

//sum := 1
//for sum < 1000 {
//	sum += sum
//}

//Forever
//If you omit the loop condition it loops forever, so an infinite loop is compactly expressed
//func main() {
//	for {
//	}
