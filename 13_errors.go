package main

import (
	"errors"
	"fmt"
)

type myError struct {
	value int
	error string
}

func (e myError) Error() string {
	return fmt.Sprintf("%d - %s", e.value, e.error)
}

func errorsExample() {
	values := []int{3, 42, 7}

	for _, v := range values {
		if r, e := work(v); e != nil {
			fmt.Println(r, e) // error
		} else {
			fmt.Println(r, e) // ok
		}
	}
}

func work(val int) (bool, error) {
	if val == 7 {
		return false, myError{val, "not a valid number"}
	}

	if val == 42 {
		return false, errors.New("42 is not a valid number")
	}

	return true, nil
}
