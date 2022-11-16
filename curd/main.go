package main

import (
	"curd/controller"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/employees", controller.GetEmployee).Methods("GET")

	r.HandleFunc("/employees/{id}", controller.GetEmployeeID).Methods("GET")

	r.HandleFunc("/employees", controller.Create).Methods("POST")

	r.HandleFunc("/employees/{id}", controller.Update).Methods("PUT")

	r.HandleFunc("/employees/{id}", controller.Delete).Methods("DELETE")

	fmt.Println("server running at 8081...")

	//starting server
	log.Fatal(http.ListenAndServe(":8081", r))

}
