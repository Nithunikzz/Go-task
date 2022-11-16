package main

import "fmt"

type rect struct {
	width  int
	height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r *rect) perm() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 5, height: 10}

	fmt.Printf("%d\n", r.area())

	fmt.Printf("%d\n", r.perm())
}
