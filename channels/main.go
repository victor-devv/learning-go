//Channels are unique to Go. They are very useful, as they are a means of sending information from one part of your program to another very easily

package main

import (
	"github.com/victor-devv/myniceprogramchannels/helpers"
	"log"
)

const numPool = 10

//take argument of intChan type channel of int
func CalculateValue(intChan chan int) {
	//generate a random number
	randomNumber := helpers.RandomNumber(numPool)

	//pass random number to channel
	intChan <- randomNumber
}

func main() {
	//create channel; a place to send info which would be received in one or more places in the program.
	//This channel can only hold integers
	intChan := make(chan int)

	//close channel when main() is done running
	defer close(intChan) //defer means whatever comes after the keyword, execute it after the current function (close) is done

	//Fire Go routine. (Run things concurrently using keyword go)
	// To make a function a routine, put the keyword GO before it
	go CalculateValue(intChan)

	//listen for response from channel
	num := <-intChan
	log.Println(num)
}




