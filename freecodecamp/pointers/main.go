package main

import (
	"fmt"
	"strings"
)

func main() {
	test("Lane", "Textio is getting better everyday!")
	test("Allan", "This pointer stuff is weird...")
	test("Tiffany", "What time will you be home for dinner?")

	messages1 := []string{
		"well shoot, this is awful",
		"dang robots",
		"dang them to heck",
	}

	messages2 := []string{
		"well shoot",
		"Allan is going straight to heck",
		"dang... that's a tough break",
	}

	test2(messages1)
	test2(messages2)
}

func removeProfanity(message *string) {
	result := *message
	result = strings.ReplaceAll(result, "dang", "****")
	result = strings.ReplaceAll(result, "shoot", "*****")
	result = strings.ReplaceAll(result, "heck", "****")
	*message = result
}

// don't touch below this line

func test2(messages []string) {
	for _, message := range messages {
		removeProfanity(&message)
		fmt.Println(message)
	}
}

type Message struct {
	Recipient string
	Text      string
}

// Don't touch above this line

func sendMessage(m Message) {
	fmt.Printf("To: %v\n", m.Recipient)
	fmt.Printf("Message: %v\n", m.Text)
}

// Don't touch below this line

func test(recipient string, text string) {
	m := Message{Recipient: recipient, Text: text}
	sendMessage(m)
	fmt.Println("=====================================")
}
