package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type ApiHandler struct {
	odd  chan int
	even chan int
}
func NewAPIHandler() *ApiHandler {
	return &ApiHandler{odd: make(chan int), even: make(chan int)}
}
func main() {
	a := ApiHandler{}
	router := mux.NewRouter()
	router.HandleFunc("/sayHello/{names}", a.SayHello)
	router.HandleFunc("/sayHello/{intvalues}", a.readDatafrom)
	http.ListenAndServe(":8080", router)
}
func (a *ApiHandler) SayHello(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	value := r.URL.Query()
	w.Write([]byte(fmt.Sprintf("%s %s", vars["names"], value["intvalues"])))
}
func (a *ApiHandler) Print(w http.ResponseWriter, r *http.Request){

	val := 0
	switch val % 2 {
	case 0:
		a.even <- val
		fmt.Println("Even")
	case 1:
		a.odd <- val
		fmt.Println("odd")
	default:
	}

}
func (a *ApiHandler) readDatafrom (){

	vars := mux.Vars(r)
	go print(vars)
	fmt.Printf("%s", vars["intvalues"])
	res := <-a.even
	fmt.Println(res)
	res = <-a.odd
	fmt.Println(res)
}
