package main

import (
	"log"
	"time"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
	BirthDate time.Time
}

func TypesAndStruct() {
	person := Person{
		FirstName: "Tega",
		LastName:  "Oke",
		Age:       29,
	}

	log.Println(person.FirstName, person.LastName, person.Age)
}
