package main

import "fmt"

func main() {

	pattren1()
	pattren2()
	pattren3()
}
func pattren1() {
	var rows int

	fmt.Print("Rows to Print the Right Angled Triangle = ")
	fmt.Scanln(&rows)

	for i := 1; i <= rows; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(" *")
		}
		fmt.Println("")
	}
}
func pattren2() {
	var rows int
	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &rows)
	for i := 0; i <= rows; i++ {
		for j := 0; j < rows-i; j++ {
			fmt.Printf(" ")
		}
		for k := 0; k < i; k++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

func pattren3() {

	n := 4
	px := n
	py := n
	for i := 1; i <= n; i++ {
		for j := 0; j <= n*2; j++ {
			if (j == px) || (j == py) {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		px = px - 1
		py = py + 1
		fmt.Print("\n")
	}
}
