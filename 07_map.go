package main

import "fmt"

func maps() {
	sampleMap := make(map[string]int)

	sampleMap["k1"] = 99
	sampleMap["k2"] = 98

	fmt.Println(sampleMap)

	delete(sampleMap, "k1")

	fmt.Println(sampleMap)

	_, ok1 := sampleMap["k1"]
	val2, ok2 := sampleMap["k2"]

	fmt.Println(ok1)
	fmt.Println(val2, ok2)
}
