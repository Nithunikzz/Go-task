package main

import (
	"log"
)

func fact(x float64) float64 {
	if x == 1 {
		return 1
	}
	return fact(x-1) * x
}

func main() {
	val := fact(5)
	log.Printf("Factorial of is: %f", val)
}
