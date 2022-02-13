//to create a package, go to terminal
//go mod init github.com/victor-devv/myniceprogram
package main

import (
	"github.com/victor-devv/myniceprogram/helpers"
	"log"
)

func main() {
	log.Println("Hello")

	var myVar helpers.SomeType

	myVar.TypeName = "Some name"
	log.Println(myVar.TypeName)
}
