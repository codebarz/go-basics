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

	log.Println("=========== JSON ===========")

}
