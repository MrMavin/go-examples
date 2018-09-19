package main

import "fmt"

// This is a constant
const PI = 3.14

func variables() {
	// Default value is empty string / 0
	var str string

	// Go will infer the type of initialized variables
	var bool = true

	// The := syntax is shorthand for declaring and initializing a variable
	integer := 10

	// Declaring and assign
	var float float32 = 5 / 3.33

	// You can assign multiple values
	var aString, aNumber, aFloatingPointNumber = "Hello", 42, PI

	fmt.Printf("%T %v\n", str, str)
	fmt.Printf("%T %v\n", bool, bool)
	fmt.Printf("%T %v\n", integer, integer)
	fmt.Printf("%T %v\n", float, float)
	fmt.Printf("%T %v", PI, PI)

	fmt.Printf("%v %v %v", aString, aNumber, aFloatingPointNumber)
}
