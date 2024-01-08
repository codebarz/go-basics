package main

import "fmt"

func concat(s1 string, s2 string) string {
	return s1 + s2
}

// when all function parameters are of the same type, the type can be declare on the last param of same type

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("startin Textio server")

	const secondsInMinute = 60
	const minuteInHour = 60
	const secondsInHour = secondsInMinute * minuteInHour

	fmt.Println("number of seconds in an hour:", secondsInHour)

	const name = "Saul Goodman"
	const openRate = 30.5

	msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent", name, openRate)

	fmt.Println(msg)

	// CONDITIONALS

	messageLen := 10
	maxMessagelen := 20
	fmt.Println("Trying to send a message of length:", messageLen)

	if messageLen <= maxMessagelen {
		fmt.Println("Message sent")
	} else {
		fmt.Println("Message not sent")
	}

	if condition := messageLen <= maxMessagelen; condition {
		fmt.Println("Message sent")
	} else {
		fmt.Println("Message not sent")
	}

	// FUNCTIONS

	fmt.Println(concat("Lane", " happy birthday!"))
	fmt.Println(concat("Elon", " hope that Tesla thing works out."))
	fmt.Println(concat("Go", " is fantastic"))
	fmt.Println(add(2, 4))

	sendSoFar := 430
	const sendToAdd = 25
	sendSoFar = incrementSends(sendSoFar, sendToAdd)
	fmt.Println("you've sent", sendSoFar, "messages")

	// Functions with multiple return values
	firstName, _ := getNames()
	fmt.Println("Welcome to Textio,", firstName)
	fmt.Println(namedReturnValues())
	fmt.Println(yearsUntilEvents(20))

	// STRUCTS

	structs()

	newTruck := truck{
		model: "TATA",
		car: car{
			wheelAmount:  4,
			headLampType: "angry",
		},
	}

	fmt.Println(newTruck.wheelAmount)

	// Define methods in struct
	rectangle := rect{
		width:  20,
		height: 10,
	}
	fmt.Println(rectangle.area())

	testAuth(authenticationInfo{
		username: "Google",
		password: "12345",
	})
	testAuth(authenticationInfo{
		username: "Bing",
		password: "98765",
	})
	testAuth(authenticationInfo{
		username: "DDG",
		password: "76921",
	})
}

type authenticationInfo struct {
	username string
	password string
}

func (authInfo authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf("Authorization: Basic %s:%s", authInfo.username, authInfo.password)
}

// don't touch below this line

func testAuth(authInfo authenticationInfo) {
	fmt.Println(authInfo.getBasicAuth())
	fmt.Println("====================================")
}

type rect struct {
	width  int
	height int
}

func (r rect) area() int {
	return r.height * r.width
}

type messageToSend struct {
	message     string
	phoneNumber int
}

type car struct {
	wheelAmount  int
	headLampType string
}

type truck struct {
	car
	model string
}

func test(m messageToSend) {
	fmt.Printf("Sending message: %s to %v\n", m.message, m.phoneNumber)
	fmt.Println("======================================================")
}

func structs() {
	test(messageToSend{message: "Thanks for signing up", phoneNumber: 122338898})
}

func incrementSends(sendSoFar, sendToAdd int) int {
	return sendSoFar + sendToAdd
}

func getNames() (string, string) {
	return "John", "Doe"
}

// named return value
func namedReturnValues() (weight, height int) {
	return
}

func yearsUntilEvents(age int) (yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental int) {
	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}

	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}

	yearsUntilCarRental = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}

	return
}
