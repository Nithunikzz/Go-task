package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type EmployeeDetails struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Position string `json:"position"`
	Salary   uint64 `json:"salary"`
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/createemp", CreateNewEmp).Methods("POST")
	fmt.Println("server running at 8000...")
	log.Fatal(http.ListenAndServe(":8080", r))

}

func CreateNewEmp(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal("error in reading body", err)
	}

	//fmt.Println(body)

	var emp EmployeeDetails
	json.Unmarshal(body, &emp)

	json.NewEncoder(w).Encode(emp)

}
