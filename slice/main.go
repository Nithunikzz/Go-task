package main

import "fmt"

func main() {
	a := make([]string, 3)
	a[0] = "b"
	a[1] = "c"
	a[2] = "d"
	fmt.Println(a)

	e := a[1:2]

	fmt.Println(e)

	fmt.Println(len(a))
}
