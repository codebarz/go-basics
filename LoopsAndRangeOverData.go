package main

import "log"

func LoopsAndRangeOverData() {
	log.Println("=========== Loops and Range Over Data ===========")

	for i := 0; i <= 3; i++ {
		log.Println("This value of i is", i)
	}

	animals := []string{"dog", "cat", "zibra", "eagle"}

	for i, animal := range animals {
		log.Println(i, animal)
	}

	type User struct {
		Firstname string
		Lastname  string
		Age       int
	}

	var users []User

	users = append(users, User{"Tega", "Mike", 30})
	users = append(users, User{"Mike", "Adenuga", 77})

	for i, user := range users {
		log.Println(i, user.Firstname, user.Lastname, user.Age)
	}

	var usersMap = make(map[string]User)

	usersMap["male"] = User{"Tega", "Mike", 30}
	usersMap["female"] = User{"Zoe", "Michel", 1}

	for key, user := range usersMap {
		log.Println(key, user)
	}

	log.Println("=========== Loops and Range Over Data ===========")

}
