package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {
	//FOR LOOP

	//for i := 0; i <= 10; i++ {
	//	log.Println(i)
	//}

	//ranging over data

	//mySlice := []string{"dog", "cat", "horse", "fish", "banana"}
	//for _, x := range mySlice{
	//	log.Println(x)
	//}

	var mySlice []User
	u1 := User{
		FirstName: "Victor",
		LastName:  "Adrien",
	}

	u2 := User{
		FirstName: "Jane",
		LastName:  "Biddy",
	}

	mySlice = append(mySlice, u1)
	mySlice = append(mySlice, u2)

	for i, x := range mySlice{
		//log.Println(i, x)
		log.Println(i, x.FirstName)
	}

	//myMap := make(map[string]string)
	//myMap["dog"] = "dog"
	//myMap["fish"] = "fish"
	//myMap["hat"] = "hat"
	//
	//for i, x := range myMap{
	//	log.Println(i, x)
	//}
}
