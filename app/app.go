package app

import (
	"log"
	"net/http"
)

func Start() {
	//Define the routes
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	//Starting the server
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
