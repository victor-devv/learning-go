package main

import "log"

func main() {
	// IF ELSE STATEMENT

	//var isTrue bool
	//
	//isTrue = true
	//
	//if isTrue == true {
	//	log.Println("isTrue is", isTrue)
	//} else {
	//	log.Println("isTrue is", isTrue)
	//}

	//cat := "cat"
	//
	//if cat == "cat" {
	//	log.Println("Cat is cat")
	//} else {
	//	log.Println("Cat is not cat")
	//}

	//myNum := 100
	//isTrue := false
	//
	//if myNum > 99 && !isTrue {
	//	log.Println("myNum is greater than 99 and isTrue is set to true")
	//} else if myNum < 100 && isTrue {
	//	log.Println(1)
	//} else if myNum == 101 || isTrue {
	//	log.Println(2)
	//} else if myNum > 1000 && isTrue == false {
	//	log.Println(3)
	//} else {
	//	log.Println("All conditions set, are not true")
	//}

	//SWITCH STATEMENTS

	animal := "cat"

	switch animal {
	case "cat":
		log.Println("animal is set to cat")

	case "dog":
		log.Println("animal is set to dog")

	case "fish":
		log.Println("animal is set to fish")

	default:
		log.Println("animal is something else")
	}
}