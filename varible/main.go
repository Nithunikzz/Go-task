package main

import (
	"fmt"
)

type names struct {
	name []string
}

func length(n names) {
	fmt.Println(n)
}

func main() {
	n := names{
		name: []string{
			"nithu",
			"alan",
			"akash",
			"ahin",
		},
	}
	length(n)
}
