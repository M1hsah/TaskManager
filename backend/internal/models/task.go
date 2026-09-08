package models

import (
	"database/sql"
	"time"
)

type Task struct {
	ID        int       `json:"id"`
	ProjectID int       `json:"projectId"`
	Title     string    `json:"title"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}

type TaskModel struct {
	DB *sql.DB
}

func (m *TaskModel) GetTasksByProject(projectId int) ([]Task, error) {
	rows, err := m.DB.Query("SELECT id, project_id, title, status, created_at FROM tasks WHERE project_id = $1", projectId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []Task{}
	for rows.Next() {
		var t Task
		err := rows.Scan(&t.ID, &t.ProjectID, &t.Title, &t.Status, &t.CreatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (m *TaskModel) Insert(title string, projectId int) (Task, error) {
	var t Task
	row := m.DB.QueryRow(
		"INSERT INTO tasks (project_id, title) VALUES ($1, $2) RETURNING id, project_id, title, status, created_at",
		projectId,
		title,
	)
	err := row.Scan(&t.ID, &t.ProjectID, &t.Title, &t.Status, &t.CreatedAt)
	if err != nil {
		return Task{}, err
	}
	return t, nil
}
