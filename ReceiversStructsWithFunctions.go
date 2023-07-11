package main

import "log"

type NewPerson struct {
	FirstName string
	LastName  string
	Age       int
}

func (p *NewPerson) printPersonName() {
	log.Println("Person first name and last name is", p.FirstName, p.LastName)
}

func (p *NewPerson) printPersonAge() {
	log.Println("Person's age is", p.Age)
}

func ReceiversStructsWithFunctions() {

	person := NewPerson{
		FirstName: "Tega",
		LastName:  "Oke",
		Age:       29,
	}

	person.printPersonName()
	person.printPersonAge()

}
