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
	// TODO 1: Create an explicit mux instead of relying on the
	// package-level http.HandleFunc (which registers on Go's global
	// http.DefaultServeMux).
	// Hint: http.NewServeMux() returns a *http.ServeMux.
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", healthHandler)
	mux.HandleFunc("GET /projects", projectsHandler)
	// TODO 2: Register both handlers on your mux, not on the package-level
	// http functions. Use *mux*.HandleFunc(...) this time.
	// Since Go 1.22, ServeMux patterns can include an HTTP method, e.g.
	// "GET /health" instead of just "/health" — try that syntax for both
	// routes so only GET requests match.

	// TODO 3: Pass your mux (not nil) as the second argument to
	// ListenAndServe so it actually gets used as the router.
	log.Fatal(http.ListenAndServe(":8080", mux))
}
