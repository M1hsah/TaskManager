package main

import (
	"encoding/json"
	"log"
	"net/http"

	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
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
	connectionString := "postgres://hashimkhalid@localhost:5432/taskmanager?sslmode=disable"
	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	log.Println("Connected to the database successfully")
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", healthHandler)
	mux.HandleFunc("GET /projects", projectsHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
