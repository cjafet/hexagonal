package app

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter()

	// define routes
	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}