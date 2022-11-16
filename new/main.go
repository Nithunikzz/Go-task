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
	width  int
	height int
}

func (r rectangle) calculateArea() int {
	return r.height * r.width
}

type triangle struct {
	edge int
}

func (t triangle) calculateArea() int {
	return t.edge * t.edge
}

type circle struct {
	radius int
}

func (c circle) calculateArea() int {
	return 3 * c.radius * c.radius
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
	fmt.Printf("s2: %v\n", s2)
	fmt.Printf("s2.p: %v\n", *s2.exPointer)
	s1 = s2
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s1.p: %v\n", *s1.exPointer)
	e = 100
	fmt.Printf("s1.p: %v\n", *s1.exPointer)
	fmt.Printf("s2.p: %v\n", *s2.exPointer)
	s10 := square{
		width:  5,
		height: 5,
	}
	r10 := rectangle{10, 10}
	t10 := triangle{15}
	c10 := circle{20}
	printArea(s10)
	printArea(r10)
	printArea(t10)
	printArea(c10)
	printType(&e)
	var i interface{}
	i = e
	
	printType(t10)

}
func printType(e interface{}) {
	switch e.(type) {
	case int:
		fmt.Println("int")
	case *int:
		fmt.Println("pointer int")
	case square:
		fmt.Println("square")
	default:
		fmt.Println("default")
	}
}
func printArea(s Shape) {
	t := ""
	switch s.(type) {
	case square:
		t = "square"
	case rectangle:
		t = "rectangle"
	case triangle:
		t = "triangle"
	case circle:
		t = "circle"
	}
	fmt.Printf("type: %v, Area: %v\n", t, s.calculateArea())
}
