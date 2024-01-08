package main

import "fmt"

func main() {
	test(10)
	test(20)
	test(30)
	test(40)
	test(50)

	test2(10.00)
	test2(20.00)
	test2(30.00)
	test2(5.1)
	test2(40.00)
	test2(50.00)

	test3(1.1, 5)
	test3(1.3, 10)
	test3(1.35, 25)

	fizzbuzz()

}

func fizzbuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
	actualCostInPennies := 1.0
	maxMessagesToSend := 1
	for actualCostInPennies <= float64(maxMessagesToSend) {
		actualCostInPennies *= costMultiplier
		maxMessagesToSend++
	}

	if actualCostInPennies > float64(maxCostInPennies) {
		maxMessagesToSend--
	}

	return maxMessagesToSend
}

// don't touch below this line

func test3(costMultiplier float64, maxCostInPennies int) {
	maxMessagesToSend := getMaxMessagesToSend(costMultiplier, maxCostInPennies)
	fmt.Printf("Multiplier: %v\n",
		costMultiplier,
	)
	fmt.Printf("Max cost: %v\n",
		maxCostInPennies,
	)
	fmt.Printf("Max messages you can send: %v\n",
		maxMessagesToSend,
	)
	fmt.Println("====================================")
}

func maxMessages(thresh float64) int {
	totalCost := 0.0
	for i := 0; ; i++ {
		totalCost += 1.0 + (0.01 * float64(i))
		if totalCost > thresh {
			return i
		}
	}
}

// don't edit below this line

func test2(thresh float64) {
	fmt.Printf("Threshold: %.2f\n", thresh)
	max := maxMessages(thresh)
	fmt.Printf("Maximum messages that can be sent: = %v\n", max)
	fmt.Println("===============================================================")
}

func bulkSend(numMessages int) float64 {
	totalFee := 1.0

	for i := 1; i < numMessages; i++ {
		totalFee += 1.0 + (0.01 * float64(i))

	}

	return totalFee
}

func test(numMessages int) {
	fmt.Printf("Sending %v messages\n", numMessages)
	cost := bulkSend(numMessages)
	fmt.Printf("Bulk send complete! Cost = %.2f\n", cost)
	fmt.Println("===============================================================")
}
