package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

func main() {

	httpServer := http.Server{
		Addr:    fmt.Sprintf(":8080"),
		Handler: handleRequest(),
	}
	err := httpServer.ListenAndServe()
	if err != nil && errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("unexpected :%v", err)
		os.Exit(1)
	}
}
func handleRequest() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/sayHello/{names}", SayHello).Methods("GET")
	return router
}

//action
func SayHello(w http.ResponseWriter, r *http.Request) {
	pathParam := strings.Split(r.URL.Path, "/")

	names := pathParam[2]

	fmt.Printf("values: %s\n", names)

}
