package main

import (
	// "fmt"
	"log"
)

//main function doesn't take any parameters
func main() {
	var whatToSay string
	var saySomethingElse string
	// var world string

	// whatToSay, world = saySomething("Hello")
	whatToSay, _ = saySomething("Hello") //_ means ignore the second returned parameter

	//fmt.Println(whatToSay)
	log.Println(whatToSay)

	saySomethingElse, _ = saySomething("Goodbye")
	log.Println(saySomethingElse)
}

func saySomething(s string) (string, string) {
	return s, "World"
}
