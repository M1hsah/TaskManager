package main

import (
	"encoding/json"
	"net/http"

	"taskmanager/internal/models"
)

type healthResponse struct {
	Status string `json:"status"`
}

type projectsResponse struct {
	Projects []models.Project `json:"projects"`
}

type createProjectInput struct {
	Name string `json:"name"`
}

func (app *application) healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := healthResponse{Status: "ok"}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (app *application) projectsHandler(w http.ResponseWriter, r *http.Request) {
	projects, err := app.projects.GetAll()
	if err != nil {
		http.Error(w, "Failed to fetch projects", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := projectsResponse{Projects: projects}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (app *application) createProjectHandler(w http.ResponseWriter, r *http.Request) {
	var input createProjectInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if input.Name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	project, err := app.projects.Insert(input.Name)
	if err != nil {
		http.Error(w, "Failed to create project", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(project)
}
