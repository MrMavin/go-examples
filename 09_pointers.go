package main

import "fmt"

func pointers() {
	i := 5

	fmt.Println(i, &i)

	toZero(i)

	fmt.Println(i)

	toZeroPtr(&i)

	fmt.Println(i)
}

func toZero(i int) {
	i = 0
}

func toZeroPtr(i *int) {
	*i = 0
}
