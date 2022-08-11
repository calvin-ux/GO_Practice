package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("My favourite number is: ", rand.Intn(50))
	fmt.Println(rand.Float64())
}
