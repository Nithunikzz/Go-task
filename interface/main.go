package main

import (
	"fmt"
)

type square struct {
	width     int
	height    int
	exPointer *int
}

func (s square) calculateArea() int {
	return s.height * s.width
}

type rectangle struct {
	width     int
	height    int
	exPointer *int
}

func (r rectangle) calculateArea() int {
	return r.height * r.width
}

type Shape interface {
	calculateArea() int
}

func main() {
	s1 := square{}
	fmt.Printf("s1: %v", s1)
	var e int = 10

	s2 := square{
		width:     5,
		height:    5,
		exPointer: &e,
	}
	s10 := square{
		width:  5,
		height: 5,
	}
	r10 := rectangle{10, 10, &e}
	printArea(s10)
	printArea(r10)

	fmt.Printf("s2: %v\n", s2)
	fmt.Printf("s2.p: %v\n", *s2.exPointer)
	s1 = s2
	fmt.Printf("s2: %v\n", s1)
	fmt.Printf("s2.p: %v\n", *s1.exPointer)
}

func printArea(s Shape) {
	t := ""
	switch s.(type) {
	case square:
		t = "square"
	case rectangle:
		t = "rectangle"

	}
	fmt.Printf("type: %v, Area: %v\n", t, s.calculateArea())
}
