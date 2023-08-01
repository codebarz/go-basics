package main

import "log"

type Animal interface {
	Say() string
	NumberOfLegs() int
}

type Dog struct {
	Breed string
	Color string
}

func interfaces() {

	log.Println("=========== INTERFACES ===========")

	dog := Dog{
		Breed: "Rot",
		Color: "black",
	}

	action(&dog)

	log.Println("=========== INTERFACES ===========")

}

func action(a Animal) {
	log.Println("The animal says ", a.Say(), " and the animal has ", a.NumberOfLegs(), " legs")
}

func (d *Dog) Say() string {
	return "Woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}
