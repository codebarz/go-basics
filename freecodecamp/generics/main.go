package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	test([]email{}, "email")
	test([]email{
		{
			"Hi Margo",
			"janet@example.com",
			"margo@example.com",
		},
		{
			"Hey Margo I really wanna chat",
			"janet@example.com",
			"margo@example.com",
		},
		{
			"ANSWER ME",
			"janet@example.com",
			"margo@example.com",
		},
	}, "email")
	test([]payment{
		{
			5,
			"jane@example.com",
			"sally@example.com",
		},
		{
			25,
			"jane@example.com",
			"mark@example.com",
		},
		{
			1,
			"jane@example.com",
			"sally@example.com",
		},
		{
			16,
			"jane@example.com",
			"margo@example.com",
		},
	}, "payment")

	test2(subscription{
		userEmail: "john@example.com",
		startDate: time.Now().UTC(),
		interval:  "yearly",
	},
		[]subscription{},
		1000.00,
	)
	test2(subscription{
		userEmail: "jane@example.com",
		startDate: time.Now().UTC(),
		interval:  "monthly",
	},
		[]subscription{
			{
				userEmail: "jane@example.com",
				startDate: time.Now().UTC().Add(-time.Hour * 24 * 7),
				interval:  "monthly",
			},
			{
				userEmail: "jane@example.com",
				startDate: time.Now().UTC().Add(-time.Hour * 24 * 7 * 52 * 2),
				interval:  "yearly",
			},
		},
		686.20,
	)
	test2(oneTimeUsagePlan{
		userEmail:        "dillon@example.com",
		numEmailsAllowed: 5000,
	},
		[]oneTimeUsagePlan{},
		756.20,
	)
	test2(oneTimeUsagePlan{
		userEmail:        "dalton@example.com",
		numEmailsAllowed: 100000,
	},
		[]oneTimeUsagePlan{
			{
				userEmail:        "dalton@example.com",
				numEmailsAllowed: 34200,
			},
		},
		32.20,
	)

	testBiller[user](
		userBiller{Plan: "basic"},
		user{UserEmail: "joe@example.com"},
	)
	testBiller[user](
		userBiller{Plan: "basic"},
		user{UserEmail: "samuel.boggs@example.com"},
	)
	testBiller[user](
		userBiller{Plan: "pro"},
		user{UserEmail: "jade.row@example.com"},
	)
	testBiller[org](
		orgBiller{Plan: "basic"},
		org{Admin: user{UserEmail: "challis.rane@example.com"}},
	)
	testBiller[org](
		orgBiller{Plan: "pro"},
		org{Admin: user{UserEmail: "challis.rane@example.com"}},
	)
}

// PARAMETRIC CONTRAINTS
// PASSING GENERIC AS PARAMS TO INTERFACE

type biller[C customer] interface {
	Charge(C) bill
	Name() string
}

// don't edit below this line

type userBiller struct {
	Plan string
}

func (ub userBiller) Charge(u user) bill {
	amount := 50.0
	if ub.Plan == "pro" {
		amount = 100.0
	}
	return bill{
		Customer: u,
		Amount:   amount,
	}
}

func (sb userBiller) Name() string {
	return fmt.Sprintf("%s user biller", sb.Plan)
}

type orgBiller struct {
	Plan string
}

func (ob orgBiller) Name() string {
	return fmt.Sprintf("%s org biller", ob.Plan)
}

func (ob orgBiller) Charge(o org) bill {
	amount := 2000.0
	if ob.Plan == "pro" {
		amount = 3000.0
	}
	return bill{
		Customer: o,
		Amount:   amount,
	}
}

type customer interface {
	GetBillingEmail() string
}

type bill struct {
	Customer customer
	Amount   float64
}

type user struct {
	UserEmail string
}

func (u user) GetBillingEmail() string {
	return u.UserEmail
}

type org struct {
	Admin user
	Name  string
}

func (o org) GetBillingEmail() string {
	return o.Admin.GetBillingEmail()
}

func testBiller[C customer](b biller[C], c C) {
	fmt.Printf("Using '%s' to create a bill for '%s'\n", b.Name(), c.GetBillingEmail())
	bill := b.Charge(c)
	fmt.Printf("Bill created for %v dollars\n", bill.Amount)
	fmt.Println(" --- ")
}

// GENERICS CONSTRAINTS
// YOU MIGHT NEED THE GENERIC TO KNOW SOMTHING ABOUT THE TYPE

func chargeForLineItem[T lineItem](newItem T, oldItems []T, balance float64) ([]T, float64, error) {
	if balance < newItem.GetCost() {
		return nil, 0.0, errors.New("insufficient funds")
	}

	oldItems = append(oldItems, newItem)

	updatedBalance := balance - newItem.GetCost()

	return oldItems, updatedBalance, nil
}

// don't edit below this line

type lineItem interface {
	GetCost() float64
	GetName() string
}

type subscription struct {
	userEmail string
	startDate time.Time
	interval  string
}

func (s subscription) GetName() string {
	return fmt.Sprintf("%s subscription", s.interval)
}

func (s subscription) GetCost() float64 {
	if s.interval == "monthly" {
		return 25.00
	}
	if s.interval == "yearly" {
		return 250.00
	}
	return 0.0
}

type oneTimeUsagePlan struct {
	userEmail        string
	numEmailsAllowed int
}

func (otup oneTimeUsagePlan) GetName() string {
	return fmt.Sprintf("one time usage plan with %v emails", otup.numEmailsAllowed)
}

func (otup oneTimeUsagePlan) GetCost() float64 {
	const costPerEmail = 0.03
	return float64(otup.numEmailsAllowed) * costPerEmail
}

func test2[T lineItem](newItem T, oldItems []T, balance float64) {
	fmt.Println(" --- ")
	fmt.Printf("Charging customer for a '%s', current balance is %v...\n", newItem.GetName(), balance)
	newItems, newBalance, err := chargeForLineItem(newItem, oldItems, balance)
	if err != nil {
		fmt.Printf("Got error: %v\n", err)
		return
	}
	fmt.Printf("New balance is: %v. Total number of line items is now %v\n", newBalance, len(newItems))
}

// GENERIC INTRO

func getLast[T any](s []T) T {
	sliceLen := len(s)
	var zero T

	if sliceLen == 0 {
		return zero
	}

	return s[sliceLen-1]
}

// don't edit below this line

type email struct {
	message        string
	senderEmail    string
	recipientEmail string
}

type payment struct {
	amount         int
	senderEmail    string
	recipientEmail string
}

func test[T any](s []T, desc string) {
	last := getLast(s)
	fmt.Printf("Getting last %v from slice of length: %v\n", desc, len(s))
	for i, v := range s {
		fmt.Printf("Item #%v: %v\n", i+1, v)
	}
	fmt.Printf("Last item in list: %v\n", last)
	fmt.Println(" --- ")
}
