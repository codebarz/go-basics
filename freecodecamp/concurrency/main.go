package main

import (
	"fmt"
	"time"
)

// go rotuine is a separate thread of execution
// you can't capture a return value from a function call in a go routine

func main() {
	test("Hello there Stacy!")
	test("Hi there John!")
	test("Hey there Jane!")
}

func sendEmail(message string) {
	go func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received: '%s'\n", message)
	}()
	fmt.Printf("Email sent: '%s'\n", message)
}

// Don't touch below this line

func test(message string) {
	sendEmail(message)
	time.Sleep(time.Millisecond * 500)
	fmt.Println("========================")
}
