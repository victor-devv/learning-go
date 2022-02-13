package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {
	// myMap := make(map[string]string)

	// myMap["dog"] = "Samson"
	// myMap["other-dog"] = "Cassie"

	// myMap["dog"] = "Fido"

	// log.Println(myMap["dog"])
	// log.Println(myMap["other-dog"])

	// myMap := make(map[string]int)

	// myMap["First"] = 1
	// myMap["Second"] = 2

	// log.Println(myMap["First"])
	// log.Println(myMap["Second"])

	myMap := make(map[string]User)

	me := User{
		FirstName: "Victor",
		LastName:  "Adrien",
	}

	myMap["me"] = me

	log.Println(myMap["me"].FirstName)

	//note that the go compiler randomizes maps. So you cant depend on the order in which you entered the map values. Depend on the keys only
}
