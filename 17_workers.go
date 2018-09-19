package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func workers() {
	jobs := make(chan string, 10)
	results := make(chan string, 10)

	for i := 1; i <= 2; i++ {
		go worker(i, jobs, results)
	}

	for j := 0; j < 10; j++ {
		jobs <- randomString()
	}

	close(jobs)

	finished := false

	for !finished {
		select {
		case result := <-results:
			fmt.Println(result)
		case <-time.After(5 * time.Second):
			finished = true
			close(results)
		}
	}

	fmt.Println("Done.")
}

func worker(id int, jobs <-chan string, results chan<- string) {
	fmt.Println("Started worker", id)
	for job := range jobs {
		fmt.Println(id, "- carrying a new job")
		time.Sleep(time.Second)
		results <- strings.ToUpper(job)
		fmt.Println(id, "- finished job")
	}
	fmt.Println("Stopped worker", id)
}

const charList = "abcdefghijklmnopqrstuvwxyz"

func randomString() string {
	str := make([]byte, 10)

	for i := 0; i < 10; i++ {
		str[i] = charList[rand.Intn(len(charList))]
	}

	return string(str)
}
