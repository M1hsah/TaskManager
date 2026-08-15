package main

import (
	"log"
	"net/http"

	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"

	"taskmanager/internal/models"
)

type application struct {
	projects *models.ProjectModel
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
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", app.healthHandler)
	mux.HandleFunc("GET /projects", app.projectsHandler)
	mux.HandleFunc("POST /projects", app.createProjectHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
