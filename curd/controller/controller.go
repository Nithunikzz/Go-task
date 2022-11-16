package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Employee struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

var employees = make(map[string]Employee)

func GetEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func GetEmployeeID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, value := range employees {
		if value.Id == params["id"] {
			json.NewEncoder(w).Encode(value)
			return

		}

	}

}

func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp Employee

	_ = json.NewDecoder(r.Body).Decode(&emp)
	//key := strconv.Itoa(rand.Intn(1000000000))
	employees[emp.Id] = emp
	w.Write([]byte("Success"))

}
func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for i, emp := range employees {
		if emp.Id == params["id"] {
			var emp Employee
			_ = json.NewDecoder(r.Body).Decode(&emp)
			//employees[i] = emp
			//delete(employees, i)
			employees[i] = emp
			w.Write([]byte("Success"))

			break
		}
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for i, emp := range employees {
		if emp.Id == params["id"] {
			employees[i] = emp
			delete(employees, i)
			w.Write([]byte("Success"))

			break
		}
	}
}
