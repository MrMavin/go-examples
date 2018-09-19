package main

import "fmt"

type points struct {
	a, b int
}

func (p points) sum() int {
	return p.a + p.b
}

func (p *points) multiply(val int) {
	p.a *= val
	p.b *= val
}

func methods() {
	points := points{2, 3}

	fmt.Println(points)
	fmt.Println(points.sum())
	points.multiply(2)
	fmt.Println(points.sum())
}
