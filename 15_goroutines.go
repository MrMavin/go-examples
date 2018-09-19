package main

import (
	"fmt"
	"time"
)

func goroutines() {
	go printString("Hello")
	go printString("World")

	fmt.Scanln()
	fmt.Println("Ok.")
}

func printString(str string) {
	for i := 0; i < 5; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(str)
	}
}
