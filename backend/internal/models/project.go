package models

import (
	"database/sql"
	"time"
)

type Project struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

type ProjectModel struct {
	DB *sql.DB
}

func (m *ProjectModel) Insert(name string) (Project, error) {
	var p Project
	row := m.DB.QueryRow(
		"INSERT INTO projects (name) VALUES ($1) RETURNING id, name, created_at",
		name,
	)
	err := row.Scan(&p.ID, &p.Name, &p.CreatedAt)
	if err != nil {
		return Project{}, err
	}
	return p, nil
}

func (m *ProjectModel) GetAll() ([]Project, error) {
	rows, err := m.DB.Query("SELECT id, name, created_at FROM projects")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	projects := []Project{}
	for rows.Next() {
		var p Project
		err := rows.Scan(&p.ID, &p.Name, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return projects, nil
}
