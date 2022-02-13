package main

import "log"

type Animal interface {
	Says() string //must have function says that returns a string
	NumberOfLegs() int //must have function NumberOfLegs that returns an integer
}

type Dog struct {
	Name string
	Breed string
}

type Gorrilla struct {
	Name string
	Color string
	NumberOfTeeth int
}

func main() {
	//Interface
	//An interface is a contract which simply states that any type that uses this interface must follow the set of rules for that interface

	dog := Dog {
		Name: "Samson",
		Breed: "Boerboel",
	}

	PrintInfo(dog)

	gorilla := Gorrilla{
		Name: "King Kong",
		Color: "Black",
		NumberOfTeeth: 32,
	}

	PrintInfo(gorilla) //this will fail if the type gorilla doesn't pass the Animal interface implemented in the PrintInfo function

}

//Assign these functions to the struct 'Dog' so that the type satisfies the interface when we call PrintInfo
func (d Dog) Says() string {
	return "Woof"
}

func (d Dog) NumberOfLegs() int {
	return 4
}

func (d Gorrilla) Says() string {
	return "Grunt"
}

func (d Gorrilla) NumberOfLegs() int {
	return 2
}

//Whatever animal we pass must follow the Animal interface rule
func PrintInfo(a Animal) {
	log.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")
}