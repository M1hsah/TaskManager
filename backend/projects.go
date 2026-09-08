package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"taskmanager/internal/models"
)

type getProjectResponse struct {
	Project models.Project `json:"project"`
}

type getProjectsResponse struct {
	Projects []models.Project `json:"projects"`
}

type createProjectRequest struct {
	Name string `json:"name"`
}

func (app *application) getProjectsHandler(w http.ResponseWriter, r *http.Request) {
	projects, err := app.projects.GetAll()
	if err != nil {
		http.Error(w, "Error: Failed to fetch projects", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := getProjectsResponse{Projects: projects}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (app *application) getProjectHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	parsedId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Error: Invalid Syntax", http.StatusBadRequest)
		return
	}
	project, err := app.projects.Get(parsedId)
	if errors.Is(err, models.ErrNoRecord) {
		http.Error(w, "Error: Project not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Error: Failed to fetch project", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	response := getProjectResponse{Project: project}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (app *application) deleteProjectHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	parsedId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Error: Invalid Syntax", http.StatusBadRequest)
		return
	}
	err = app.projects.Delete(parsedId)
	if errors.Is(err, models.ErrNoRecord) {
		http.Error(w, "Error: Project not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Error: Failed to delete project", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (app *application) createProjectHandler(w http.ResponseWriter, r *http.Request) {
	var input createProjectRequest
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Error: Invalid request body", http.StatusBadRequest)
		return
	}

	if input.Name == "" {
		http.Error(w, "Error: Name is required", http.StatusBadRequest)
		return
	}

	project, err := app.projects.Insert(input.Name)
	if err != nil {
		http.Error(w, "Error: Failed to create project", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(project)
}

func (app *application) updateProjectHandler(w http.ResponseWriter, r *http.Request) {
	var input createProjectRequest
	id := r.PathValue("id")
	parsedId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Error: Invalid Syntax", http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Error: Invalid request body", http.StatusBadRequest)
		return
	}

	if input.Name == "" {
		http.Error(w, "Error: Name is required", http.StatusBadRequest)
		return
	}

	project, err := app.projects.Update(input.Name, parsedId)
	if errors.Is(err, models.ErrNoRecord) {
		http.Error(w, "Error: Project not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Error: Failed to Update project", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(project)
}
