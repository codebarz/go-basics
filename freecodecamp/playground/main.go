package main

import "fmt"

type person struct {
	first_name string
	last_name  string
}

type PersonAbilities interface {
	walk() string
	talk() string
}

func main() {

	aPerson := person{
		first_name: "Tega",
		last_name:  "Oke",
	}

	fmt.Println(aPerson.walk())
	fmt.Println(aPerson.talk())
	fmt.Println(aPerson.first_name)
}

func (s person) walk() string {
	s.first_name = "Michael"
	ret := fmt.Sprintf("%s is walking", s.first_name)
	return ret
}

func (s person) talk() string {
	s.first_name = "Michael"
	ret := fmt.Sprintf("%s is talking", s.first_name)
	return ret
}
