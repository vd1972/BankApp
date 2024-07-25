package app

import (
	"log"
	"net/http"
)

func Start() {
	mux := http.NewServeMux()

	//Define the routes
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", getAllCustomers)

	//Starting the server
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
