package main

import (
	"fmt"
)

func main() {
	var a = [3]int{1, 2, 3}
	b := [5]int{4, 5, 6, 7, 8}
	c := []string{"nithu"}
	fmt.Println(c)
	c = append(c, "lijo", "derlin")
	fmt.Println(a, c)
	fmt.Println(b)
	fmt.Println(len(c))
	fmt.Println("Capacity: ", cap(b))
	d := append(c, "hi")
	fmt.Println(d)
	for _, v := range d[0:2] {
		fmt.Println(v)
	}
	for _, v := range d[:] {
		fmt.Println(v)
	}
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	// Original slice
	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("length = %d\n", len(numbers))
	fmt.Printf("capacity = %d\n", cap(numbers))

	// Create copy with only needed numbers
	neededNumbers := numbers[:len(numbers)-10]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)

	fmt.Printf("numbersCopy = %v\n", numbersCopy)
	fmt.Printf("length = %d\n", len(numbersCopy))
	fmt.Printf("capacity = %d\n", cap(numbersCopy))
}
