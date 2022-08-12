package main

import "fmt"

//empty interface can be anything
func Anything(anything interface{}) { //if you don't know the Type declare it as an interface during compile
	fmt.Println(anything)
}

func main() {
	fmt.Println("vim-go")
	Anything(3.16)
	Anything("Calvin")
	Anything(struct{}{}) //empty struct

	//key is string for map; value can be anything
	mymap := make(map[string]interface{}) //using interface here where values is anything
	mymap["name"] = "Calvinson"
	mymap["age"] = 60
	fmt.Println(mymap)
}

/*
delcaration of struct:
	type WordPlay struct{}

instantisation of struct:
	struct{Letter string} {"Sudoku"}
*/
