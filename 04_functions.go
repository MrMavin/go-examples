package main

import "fmt"

func functions() {
	fmt.Println("Classic:", classic(10, 20, 30))

	hello, world := multipleReturn()
	fmt.Println("Multiple return:", hello, world)

	x, y := nakedReturn()
	fmt.Println("Naked return:", x, y)

	sum := variadic(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(sum)

	closure := func() int {
		return 10
	}

	fmt.Println(closure())

	fmt.Println(recursive(10))
}

func classic(a, b, c int) int {
	return a + b + c
}

func multipleReturn() (string, string) {
	return "Hello", "World!"
}

func nakedReturn() (x int, y int) {
	x = 10
	y = 12
	return
}

func variadic(nums ...int) int {
	var sum int = 0

	for _, value := range nums {
		sum += value
	}

	return sum
}

func recursive(n int) int {
	if n <= 0 {
		return 1
	}

	return n * recursive(n-1)
}
