package main

import (
	"fmt"
	"math/rand"
	"time"
)

var debug = false

func channels() {
	debug = false

	done := make(chan bool)

	go simpleChannels(done)
	go bufferedChannels(done)
	go closedRangeChannel(done)

	for i := 0; i < 3; i++ {
		select {
		case <-done:
			fmt.Println("A job has been completed")
		// Timeout example, not really useful here.
		case <-time.After(10 * time.Second):
			fmt.Println("Timeout")
		}
	}
}

func simpleChannels(c chan bool) {
	if debug {
		fmt.Println("Simple is started")
	}

	messages := make(chan string)

	// We must use a goroutine because the channel isn't
	// buffered. This means that every write or read is
	// always blocking
	go addMessage("Simple First", messages)
	go addMessage("Simple Second", messages)
	go addMessage("Simple Third", messages)

	printMessage(messages, "SC")
	printMessage(messages, "SC")
	printMessage(messages, "SC")

	close(messages)

	printMessage(messages, "SC")

	c <- true
}

func bufferedChannels(c chan bool) {
	if debug {
		fmt.Println("Buffered is started")
	}

	messages := make(chan string, 2)

	// Here we could decide to not use the goroutine
	// because the channel will start to block
	// write operations after that
	// it will be full
	addMessage("Buffered First", messages)
	addMessage("Buffered Second", messages)

	printMessage(messages, "BC")

	addMessage("Buffered Third", messages) // Full again

	printMessage(messages, "BC")
	printMessage(messages, "BC")

	close(messages)

	printMessage(messages, "BC")

	c <- true
}

func closedRangeChannel(c chan bool) {
	if debug {
		fmt.Println("Close Range Channel is started")
	}

	messages := make(chan string, 3)

	addMessage("One", messages)
	addMessage("Two", messages)
	addMessage("Three", messages)

	close(messages)

	// printMessage(messages, "CRC")
	// printMessage(messages, "CRC")
	// printMessage(messages, "CRC")
	// printMessage(messages, "CRC")

	for v := range messages {
		fmt.Println("[CRC]", v)
	}

	c <- true
}

// Sending channel
func addMessage(message string, channel chan<- string) {
	sleepTime := time.Duration(rand.Intn(3000)) * time.Millisecond

	if debug {
		fmt.Printf("Received %s, waiting %d\n", message, sleepTime)
	}
	time.Sleep(sleepTime)

	channel <- message
}

// Receiving channel
func printMessage(channel <-chan string, channelName string) {
	v, ok := <-channel

	if !ok {
		fmt.Printf("[%s]Channel is closed.\n", channelName)
	} else {
		fmt.Printf("[%s]Received: %s\n", channelName, v)
	}
}
