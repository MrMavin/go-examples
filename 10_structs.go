package main

import "fmt"

type Person struct {
	name string
	age  int
}

func structs() {
	fmt.Println(Person{"Sean", 30})
	fmt.Println(Person{name: "Sean", age: 30})
	fmt.Println(Person{age: 30})

	person := Person{name: "Sean", age: 30}

	fmt.Println(person.name, person.age)
}
