package main

import "fmt"

func ranges() {
	arrayExample := [4]int{40, 41, 42, 43}

	for key, value := range arrayExample {
		fmt.Println(key, value)
	}

	mapExample := map[string]bool{"water": false, "eggs": false, "pineapplePizza": true}

	for keys := range mapExample {
		fmt.Println(keys)
	}
}
