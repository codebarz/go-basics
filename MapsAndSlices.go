package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

// Maps are immutable

func MapsAndSlices() {
	newMap := make(map[string]string)

	newMap["name"] = "Tega"

	log.Println(newMap["name"])

	userMap := make(map[string]User)

	user := User{
		FirstName: "Tega",
		LastName:  "Mike",
	}

	userMap["user"] = user

	log.Println(userMap["user"].LastName)
}
