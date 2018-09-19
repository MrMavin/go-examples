package main

import (
	"fmt"
	"math/rand"
	"time"
)

func channels() {
	go simpleChannels()
	go bufferedChannels()

	fmt.Scanln()
}

func simpleChannels() {
	messages := make(chan string)

	// We must use a goroutine because the channel isn't
	// buffered. This means that every write or read is
	// always blocking
	go addMessage("First", messages)
	go addMessage("Second", messages)
	go addMessage("Third", messages)

	fmt.Println("Received ", <-messages)
	fmt.Println("Received ", <-messages)
	fmt.Println("Received ", <-messages)
}

func addMessage(message string, channel chan string) {
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	channel <- message
}

func bufferedChannels() {
	messages := make(chan string, 2)

	// Here we could decide to not use the goroutine
	// because the channel will start to block
	// write operations after that
	// it will be full
	messages <- "One"
	messages <- "Two" // Full

	fmt.Println(<-messages) // Emptied by one item

	messages <- "Three" // Full again

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
