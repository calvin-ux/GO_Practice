package main

import "fmt"

//Interface helps  in enforcing style
//def set of functions for class to follow like classes and obj
//like game character Walk(), Talk() ---> Follow interface
//like game NPC Walk(), Talk() ---> Follow interface
type Car interface {
	Drive()
	Stop()
}

//Lambo
type Lambo struct {
	LamboModel string
}

//enforceing interface
//for function to follow instruction
func NewModel(arg string) Car { //a way to make sure Struct use the interface e.g. both Drive(), and Stop()
	return &Lambo{arg}
}

//Cheay
type Chevy struct {
	ChevyModel string
}

func (l *Lambo) Stop() {
	fmt.Println("Stopping Lambo") //instead of returning type this returns interface e.g. Car
}

func (l *Lambo) Drive() {
	fmt.Println("Lambo on the move") // returns type
	fmt.Println(l.LamboModel)
}

func (c *Chevy) Drive() {
	fmt.Println("Chevy on the move")
	fmt.Println(c.ChevyModel)
}

func main() {
	l := Lambo{"Gallardo"}
	c := Chevy{"C300"}
	l.Drive()
	c.Drive()
}
