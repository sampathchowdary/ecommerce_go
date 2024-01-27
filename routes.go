package main

import (
	"github.com/gorilla/mux"
)

func addRoutes(router *mux.Router) {

	router.HandleFunc("/products", controller.getProducts).Methods("GET")

}
