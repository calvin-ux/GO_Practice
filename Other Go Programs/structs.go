package main

import "fmt"

type Car struct {
	Name    string
	Years   int
	ModelNo int
}

//using methods with struct(ures) way
func (e Car) Print() {
	fmt.Println(e)
}
func (e Car) Drive() {
	fmt.Println("Vhroom!! Vhroom!!")
}
func (e Car) GetName() string {
	return e.Name
}

func main() {
	//1st way
	c := Car{"Chevy", 1, 209108} //initialise struct as objects
	fmt.Println(c)
	//2nd way for large data structures
	v := Car{
		Name:    "Viper",
		Years:   5,
		ModelNo: 5000,
	}
	fmt.Println(v)

	//3rd way
	e := Car{
		Name:    "Spartan",
		Years:   1200,
		ModelNo: 300,
	}
	//call function
	e.Print()
	e.Drive()
	fmt.Println(e.GetName())
}
