package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prtha112/go_tha_test_action/handlers"
)

func main() {
	r := mux.NewRouter()

	healthHandler := &handlers.HealthHandler{}
	r.Handle("/health", healthHandler).Methods("GET")

	log.Println("Server started on port 8095")
	log.Fatal(http.ListenAndServe(":8095", r))
}
