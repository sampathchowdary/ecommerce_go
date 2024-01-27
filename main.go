package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Print("hello api")
	fmt.Fprintf(w, "hello \n")
}

func main() {

	router := mux.NewRouter()
	controllerFunc := controller
	fmt.Print("sampath")
	router.HandleFunc("/hello", hello)
	// addRoutes(router)
	router.HandleFunc("/products", controller.getProducts).Methods("GET")
	port := 8080
	fmt.Printf("Server started on port %d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}
