package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type healthResponse struct {
	Status string `json:"status"`
}

type projectsResponse struct {
	Message string `json:"message"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := healthResponse{Status: "ok"}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func projectsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := projectsResponse{Message: "Not implemented yet"}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", healthHandler)
	mux.HandleFunc("GET /projects", projectsHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
