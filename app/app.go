package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	//mux := http.NewServeMux()
	router := mux.NewRouter()

	//wiring
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerService())}

	//Define the routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	//Starting the server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
