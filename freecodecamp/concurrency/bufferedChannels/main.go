package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	test("Hello John, tell Kathy I said hi", "Whazzup bruther")
	test("I find that hard to believe.", "When? I don't know if I can", "What time are you thinking?")
	test("She says hi!", "Yeah its tomorrow. So we're good.", "Cool see you then!", "Bye!")

	test2(3)
	test2(4)
	test2(5)
	test2(6)

	test3(10)
	test3(5)
	test3(20)
	test3(13)

	rand.Seed(0)
	test4(
		[]string{
			"hi friend",
			"What's going on?",
			"Welcome to the business",
			"I'll pay you to be my friend",
		},
		[]string{
			"Will you make your appointment?",
			"Let's be friends",
			"What are you doing?",
			"I can't believe you've done this.",
		},
	)
	test4(
		[]string{
			"this song slaps hard",
			"yooo hoooo",
			"i'm a big fan",
		},
		[]string{
			"What do you think of this song?",
			"I hate this band",
			"Can you believe this song?",
		},
	)
}

func logMessages(chEmails, chSms chan string) {
	for {
		select {
		case sms, ok := <-chSms:
			if !ok {
				return
			}
			logSms(sms)
		case email, ok := <-chEmails:
			if !ok {
				return
			}
			logEmail(email)
		}
	}
}

// don't touch below this line

func logSms(sms string) {
	fmt.Println("SMS:", sms)
}

func logEmail(email string) {
	fmt.Println("Email:", email)
}

func test4(sms []string, emails []string) {
	fmt.Println("Starting...")

	chSms, chEmails := sendToLogger(sms, emails)

	logMessages(chEmails, chSms)
	fmt.Println("===============================")
}

func sendToLogger(sms, emails []string) (chSms, chEmails chan string) {
	chSms = make(chan string)
	chEmails = make(chan string)
	go func() {
		for i := 0; i < len(sms) && i < len(emails); i++ {
			done := make(chan struct{})
			s := sms[i]
			e := emails[i]
			t1 := time.Millisecond * time.Duration(rand.Intn(1000))
			t2 := time.Millisecond * time.Duration(rand.Intn(1000))
			go func() {
				time.Sleep(t1)
				chSms <- s
				done <- struct{}{}
			}()
			go func() {
				time.Sleep(t2)
				chEmails <- e
				done <- struct{}{}
			}()
			<-done
			<-done
			time.Sleep(10 * time.Millisecond)
		}
		close(chSms)
		close(chEmails)
	}()
	return chSms, chEmails
}

func concurrentFib(n int) {
	intChan := make(chan int)
	go fibonacci(n, intChan)

	for number := range intChan {
		fmt.Println(number)
	}
}

func test3(n int) {
	fmt.Printf("Printing %v numbers...\n", n)
	concurrentFib(n)
	fmt.Println("==============================")
}

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
		time.Sleep(time.Millisecond * 10)
	}
	close(ch)
}

func countReports(numSentCh chan int) int {
	total := 0

	for {
		sent, ok := <-numSentCh

		if !ok {
			break
		}
		total += sent
	}

	return total
}

// don't touch below this line

func test2(numBatches int) {
	numSentCh := make(chan int)
	go sendReports(numBatches, numSentCh)

	fmt.Println("Start counting...")
	numReports := countReports(numSentCh)
	fmt.Printf("%v reports sent!\n", numReports)
	fmt.Println("========================")
}

func addEmailsToQueue(emails []string) chan string {
	emailsToSend := make(chan string, len(emails))
	for _, email := range emails {
		emailsToSend <- email
	}
	return emailsToSend
}

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
		fmt.Printf("Sent batch of %v reports\n", numReports)
		time.Sleep(time.Millisecond * 100)
	}
	close(ch)
}

// TEST SUITE - Don't Touch Below This Line

func sendEmails(batchSize int, ch chan string) {
	for i := 0; i < batchSize; i++ {
		email := <-ch
		fmt.Println("Sending email:", email)
	}
}

func test(emails ...string) {
	fmt.Printf("Adding %v emails to queue...\n", len(emails))
	ch := addEmailsToQueue(emails)
	fmt.Println("Sending emails...")
	sendEmails(len(emails), ch)
	fmt.Println("==========================================")
}
