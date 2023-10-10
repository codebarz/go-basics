package main

import (
	"log"
	"math/rand"
	"time"
)

func Channels() {
	log.Println("============= CHANNELS =============")

	intChan := make(chan int)

	defer close(intChan)

	go calculateRandom(intChan)

	num := <-intChan

	log.Println(num)

	log.Println("============= CHANNELS =============")
}

func calculateRandom(randNumber chan int) {
	random := randomNumber(200)

	randNumber <- random
}

func randomNumber(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}
