package main

import "log"

type Dog struct {
	Name string
}

func interfaces() {
	dog := Dog{"Lizzy"}
	dog.bark()
}

func (*Dog) bark() {
	log.Println("Woof woof")
}
