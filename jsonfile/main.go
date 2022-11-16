package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type user struct {
		Fruit string
		Size  string
		Color string
	}
	value := &user{Fruit: "Apple", Size: "Large", Color: "Red"}
	s1, _ := json.Marshal(value)
	fmt.Println(string(s1))
	s2 := `{"Fruit": "Apple", "Size": "Large", "Color": "Red"}`
	value2 := &user{}
	json.Unmarshal([]byte(s2), value2)
	fmt.Println(value2.Fruit, value2.Size, value2.Color)
}
