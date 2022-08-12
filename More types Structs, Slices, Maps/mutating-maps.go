package main

import "fmt"

//https://go.dev/tour/moretypes/22
//key is Answer

func main() {
	m := make(map[string]int)
	//Insert or update an element in map m:
	m["Answer"] = 42
	fmt.Println("The Value: ", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The Value: ", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])
	//Test that a key is present with a two-value assignment:
	v, ok := m["Answer"]
	fmt.Println("The Value: ", v, "Present?", ok)

}
