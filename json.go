package main

import (
	"encoding/json"
	"log"
)

type NewUser struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsBlack bool   `json:"is_black"`
}

func Json() {
	log.Println("=========== JSON ===========")
	// read json
	newJson := `
	[
		{
			"name" : "Mike",
			"age": 20,
			"is_black": false
		}
	]
	`

	var unmarshaledJson []NewUser

	err := json.Unmarshal([]byte(newJson), &unmarshaledJson)

	if err != nil {
		log.Printf("An error occured: %v", err)
	}

	log.Println(unmarshaledJson)

	//unmarshal JSON

	var userArray []NewUser

	userOne := NewUser{Name: "Tega", Age: 28, IsBlack: true}
	userTwo := NewUser{Name: "Tosin", Age: 30, IsBlack: false}

	userArray = append(userArray, userOne)
	userArray = append(userArray, userTwo)

	theJson, err := json.MarshalIndent(userArray, "", "    ")

	if err != nil {
		log.Printf("An error occured: %v", err)
	}

	// log.Println(theJson)
	log.Println(string(theJson))

	log.Println("=========== JSON ===========")

}
