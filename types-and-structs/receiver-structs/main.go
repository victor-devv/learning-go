package main

import (
	"log"
)

type myStruct struct {
	FirstName string
}

//assign function to the struct
//first parenthesis is called a receiver and it ties the function to the struct by the pointer
//in essence, the function below is part of the struct
func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main() {
	var myVar myStruct
	myVar.FirstName = "John"

	myVar2 := myStruct{
		FirstName: "Mary",
	}

	log.Println("myVar is set to", myVar.printFirstName())
	log.Println("myVar2 is set to", myVar2.printFirstName())
}
