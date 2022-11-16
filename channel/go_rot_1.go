package main

import (
	"fmt"
	"time"
)

func writeOne() {
	for i := 0; i <= 5; i++ {
		fmt.Println("one")
		time.Sleep(1 * time.Second)
	}
}

func writeTwo() {
	for i := 0; i <= 5; i++ {
		fmt.Println("two")
		time.Sleep(2 * time.Second)
	}
}

func main() {
	go writeOne()
	go writeTwo()
	time.Sleep(10 * time.Second)
	fmt.Println("main")
}
