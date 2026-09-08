package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"taskmanager/internal/models"
)

type getTasksResponse struct {
	Tasks []models.Task `json:"tasks"`
}

type createTaskRequest struct {
	Title *string `json:"title"`
}

func (app *application) getTasksHandler(w http.ResponseWriter, r *http.Request) {
	projectId := r.PathValue("projectId")
	parsedId, err := strconv.Atoi(projectId)
	if err != nil {
		http.Error(w, "Error: Invalid Syntax", http.StatusBadRequest)
		return
	}
	tasks, err := app.tasks.GetTasksByProject(parsedId)
	if errors.Is(err, models.ErrNoRecord) {
		http.Error(w, "Error: No Tasks Associated with Project Found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Error: Failed to fetch project", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	response := getTasksResponse{Tasks: tasks}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (app *application) createTaskHandler(w http.ResponseWriter, r *http.Request) {
	var input createTaskRequest
	projectId := r.PathValue("projectId")
	parsedId, err := strconv.Atoi(projectId)
	if err != nil {
		http.Error(w, "Error: Invalid Syntax", http.StatusBadRequest)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Error: Invalid request body", http.StatusBadRequest)
		return
	}

	if input.Title == nil {
		http.Error(w, "Error: Title is required", http.StatusBadRequest)
		return
	}

	task, err := app.tasks.Insert(*input.Title, parsedId)
	if err != nil {
		http.Error(w, "Error: Failed to create task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}
