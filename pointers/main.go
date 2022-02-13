package main

import (
	// "fmt"
	"log"
)

func main() {
	var myString string

	myString = "Green"

	log.Println("myString is set to", myString)

	changeUsingPointer(&myString) //here you pass a reference to the variable you want the pointer to point to by prepending &
	log.Println("after func call, myString is set to", myString)
}

//pointers point to a specific location in memory and gives you access to it
func changeUsingPointer(s *string) {
	newValue := "Red" //short declaration. This can only be done in a function

	*s = newValue //point to the variable and assign a new value
}
