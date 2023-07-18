package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName  string
}

// Maps are immutable

func MapsAndSlices() {

	maps()

	slices()

}

func maps() {
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

func slices() {
	var newSlice []string

	newSlice = append(newSlice, "Hello")
	newSlice = append(newSlice, "World")

	log.Println(newSlice)

	anotherSlice := []int{1, 2, 3, 4, 5, 7, 6, 8, 9}

	sort.Ints(anotherSlice)

	log.Println(anotherSlice)

	log.Println(anotherSlice[5:9])

}
