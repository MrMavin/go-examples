package main

import (
	"fmt"
	"math/rand"
	"time"
)

func flow() {
	ifStatement()
	forLoop()
	whileLoop()
	switchStatement()
	deferStatement()
}

func ifStatement() {
	if i := rand.Intn(10); i >= 5 {
		fmt.Println("You were lucky!", i)
	} else {
		fmt.Println("You were unlucky...", i)
	}
}

func forLoop() {
	for i := 1; i < 4; i++ {
		fmt.Printf("%v ", i)
	}

	fmt.Println()

	j := 3

	for j < 99 {
		j += 1

		if j == 5 {
			continue
		}

		if j == 8 {
			break
		}

		fmt.Printf("%v ", j)
	}

	fmt.Println()
}

func whileLoop() {
	condition, i := true, 0

	for condition {
		i += 1

		if i >= 10 {
			condition = false
			fmt.Printf("%v %v", i, condition)
		}
	}

	fmt.Println()
}

func switchStatement() {
	weekday := time.Now().Weekday()

	switch weekday {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend :)")
	default:
		fmt.Println("It's not the weekend :(")
	}
}

func deferStatement() {
	defer fmt.Print("World!")
	defer fmt.Println("Last in... first out... It's a stack!")
	fmt.Println("Hello")
}
