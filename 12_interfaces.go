package main

import "fmt"

type vehicle interface {
	getWheelCount() int
}

type car struct{}

type scooter struct{}

func (c car) getWheelCount() int {
	return 4
}

func (s scooter) getWheelCount() int {
	return 2
}

func printWheelCount(v vehicle) {
	fmt.Println(v.getWheelCount())
}

func interfaces() {
	car := car{}
	scooter := scooter{}

	printWheelCount(car)
	printWheelCount(scooter)
}
