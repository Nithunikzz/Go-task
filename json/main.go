package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/connect2naga/logger/logging"

	"github.com/gorilla/mux"
)

type EmployeeDetails struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Locations string `json:"locations"`
}

type EndpointHandler struct {
	logger          logging.Logger
	EmployeeDetails map[string]EmployeeDetails
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/createemp", GetAllEmployees).Methods("GET")
	r.HandleFunc("/createemp", CreateNewEmp).Methods("POST")
	r.HandleFunc("/createemp/{id}", GetAllEmployeesID).Methods("GET")
	fmt.Println("server running at 8080...")
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
	fmt.Println(emp)

}
func GetAllEmployees(w http.ResponseWriter, r *http.Request) {

	var emp EmployeeDetails
	data, err := json.Marshal(emp)
	if err != nil {
		fmt.Printf("failed to marshl...")
		w.Write([]byte("error while fetching data"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
	w.WriteHeader(http.StatusOK)

}
func GetAllEmployeesID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	empId := vars["id"]

	data, err := json.Marshal(empId)
	if err != nil {
		fmt.Printf("failed to marshl...")
		w.Write([]byte("error while marshling data"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
	w.WriteHeader(http.StatusOK)
}
