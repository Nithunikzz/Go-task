package main

import "fmt"

func main() {

	a := 20
	var b *int

	b = &a
	fmt.Printf("%x\n", &a)
	fmt.Printf("%x\n", b)
	fmt.Printf("%x", a)

}
