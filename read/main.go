package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

b, err := os.ReadFile("sample.txt")
	if err != nil {
		// log.Fatalf("failed to read file: %v", err)
		log.Printf("failed to read file: %v", err)
		os.Exit(1)
	}
	fmt.Println(string(b))

	fileInfo, err := os.Stat("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nfileInfo: %+v\n", fileInfo)
}
