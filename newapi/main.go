package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/home", sayHello)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	fmt.Println("Started")
	handleRequests()
}
