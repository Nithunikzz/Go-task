package main

import "fmt"

type concat interface {
	concatination() string
}
type chars struct {
	char1 string
	char2 string
}

func (c chars) concatination() string {
	return c.char1 + c.char2
}
func value(a concat) string {
	return a.concatination()
}
func main() {
	str1 := chars{char1: "L.", char2: "Nithu"}
	fmt.Println(value(str1))
}
