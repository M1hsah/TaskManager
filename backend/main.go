package main

import (
	"encoding/json"
	"log"
	"net/http"

	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"

	"taskmanager/internal/models"
)

type application struct {
	projects *models.ProjectModel
	tasks    *models.TaskModel
}

type healthResponse struct {
	Status string `json:"status"`
}

func (app *application) healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := healthResponse{Status: "ok"}

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

	app := &application{
		projects: &models.ProjectModel{DB: db},
		tasks:    &models.TaskModel{DB: db},
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", app.healthHandler)
	mux.HandleFunc("GET /projects", app.getProjectsHandler)
	mux.HandleFunc("GET /projects/{id}", app.getProjectHandler)
	mux.HandleFunc("POST /projects", app.createProjectHandler)
	mux.HandleFunc("PATCH /projects/{id}", app.updateProjectHandler)
	mux.HandleFunc("DELETE /projects/{id}", app.deleteProjectHandler)
	mux.HandleFunc("GET /projects/{projectId}/tasks", app.getTasksHandler)
	mux.HandleFunc("POST /projects/{projectId}/tasks", app.createTaskHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
