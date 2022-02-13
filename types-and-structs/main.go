package main

import (
	"log"
	"time"
)

//Note that whenever you declare a function or a object properties as seen below, beginning with a capital letter, it will be available outside the package, else it wont

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName:   "Trevor",
		LastName:    "Sawler",
		PhoneNumber: "1 555 555-1212",
	}

	log.Println(user.FirstName, user.LastName, user.BirthDate)
}
