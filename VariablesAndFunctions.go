package main

import (
	"fmt"
	"log"
)

func Variables() {
	var aString string = "Learning Go"
	var i int = 20

	log.Println(aString)
	log.Println(i)

	somethingSaid := saySomething()

	stringFromMulti, intFromMulti := returnsMultiStrings()

	fmt.Println(somethingSaid)

	fmt.Println(stringFromMulti, intFromMulti)
}

func saySomething() string {
	return "saying something from variables and functions"
}

// returns multiple values of diff types

func returnsMultiStrings() (string, int) {
	return "string from multi", 90
}
