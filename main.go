package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type HealthHandler struct{}

func (h *HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	r := mux.NewRouter()

	healthHandler := &HealthHandler{}
	r.Handle("/health", healthHandler).Methods("GET")

	log.Println("Server started on port 8095")
	log.Fatal(http.ListenAndServe(":8095", r))
}
