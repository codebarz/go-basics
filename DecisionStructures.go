package main

import "log"

func DecisionStructures() {
	var isTrue bool = false

	if isTrue {
		log.Println("Its a truthy")
	} else if !isTrue {
		log.Println("Its a false")
	} else {
		log.Println("Don't know")
	}

	animal := "dog"

	switch animal {
	case "dog":
		log.Println("Its a dog")
	case "cat":
		log.Println("Its a cat")
	default:
		log.Println("Na sha animal")
	}
}
