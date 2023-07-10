package main

import "log"

func Pointers() {
	newString := "string value"

	log.Println(newString)
	changeStrignValueUsingPointer(&newString)
	log.Println(newString)
}

func changeStrignValueUsingPointer(s *string) {
	pointerValue := "value from pointer"

	*s = pointerValue
}
