package main

import (
	"fmt"
	"time"
)

func main() {
	// newFunction1()
	newFunction2()
}

func newFunction1() {
	i := 0
	go func() {
		fmt.Printf("one, i:%d\n", i)
		i++
		time.Sleep(1 * time.Second)
	}()

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Printf("two, i:%d\n", i)
		i++
	}()
	time.Sleep(10 * time.Second)
	fmt.Printf("newFunction1: i %d\n", i)
}

func newFunction2() {
	i := 0
	go func(i int) {
		fmt.Printf("one, i:%d\n", i)
		i++
		time.Sleep(1 * time.Second)
	}(i)

	go func(i *int) {
		fmt.Printf("one, i:%d\n", *i)
		*i++
		time.Sleep(1 * time.Second)
	}(&i)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Printf("two, i:%d\n", i)
		i++
	}()
	time.Sleep(10 * time.Second)
	fmt.Printf("newFunction2: i %d\n", i)
}
