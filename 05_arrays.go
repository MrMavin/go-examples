package main

import "fmt"

func arrays() {
	var arr [5]int

	fmt.Printf("%T %v\n", arr, arr)

	arr[4] = 100

	fmt.Printf("%T %v\n", arr[4], arr[4])

	hello := [5]string{"H", "e", "l", "l", "o"}

	fmt.Printf("%v (%v)\n", hello, len(hello))

	multi := [2][2]int{
		{1, 2},
		{3, 4},
	}

	fmt.Printf("%v", multi)
}
