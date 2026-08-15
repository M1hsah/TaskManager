package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type application struct {
	DB *sql.DB
}

type healthResponse struct {
	Status string `json:"status"`
}

type projectsResponse struct {
	Projects []project `json:"projects"`
}

type project struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

func (app *application) healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := healthResponse{Status: "ok"}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (app *application) projectsHandler(w http.ResponseWriter, r *http.Request) {
	// Query the database to fetch projects
	rows, err := app.DB.Query("SELECT id, name, created_at FROM projects")
	if err != nil {
		http.Error(w, "Failed to fetch projects", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Map the rows to the projects slice
	projects := []project{}
	for rows.Next() {
		var p project
		err = rows.Scan(&p.ID, &p.Name, &p.CreatedAt)
		if err != nil {
			http.Error(w, "Failed to scan project", http.StatusInternalServerError)
			return
		}
		projects = append(projects, p)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, "Error iterating over projects", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := projectsResponse{Projects: projects}

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

	app := &application{DB: db}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", app.healthHandler)
	mux.HandleFunc("GET /projects", app.projectsHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
