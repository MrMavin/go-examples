package main

import "fmt"

func slices() {
	helloSlice := make([]string, 5, 5)

	fmt.Printf("%T %v %v %v\n", helloSlice, helloSlice, len(helloSlice), cap(helloSlice))

	helloSlice[0] = "H"
	helloSlice[1] = "e"
	helloSlice[2] = "l"
	helloSlice[3] = "l"
	helloSlice[4] = "o"

	fmt.Println(helloSlice)
	fmt.Println(len(helloSlice))

	helloSlice = append(helloSlice, "W", "o", "r", "l", "d")
	fmt.Println(helloSlice, cap(helloSlice))

	helloCopy := make([]string, len(helloSlice))
	copy(helloCopy, helloSlice)

	fmt.Println(helloCopy)

	hello := helloSlice[:5]
	world := helloSlice[5:] // helloSlice[5:10]

	fmt.Println(hello, world)

	twoD := make([][]int, 4)
	twoD = append(twoD, []int{})

	for i := 0; i < len(twoD); i++ {
		length := i + 1
		twoD[i] = make([]int, length)
		for j := 0; j < length; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println(twoD)
}
