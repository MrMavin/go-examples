package main

import "fmt"

type info struct {
	name, surname string
	age           int
}

func (i info) String() string {
	return fmt.Sprintf("%s %s (%d)", i.surname, i.name, i.age)
}

func stringer() {
	fmt.Println(info{"John", "Doe", 22})
}
