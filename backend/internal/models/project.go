package models

import (
	"database/sql"
	"errors"
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

var ErrNoRecord = errors.New("models: no matching record found")

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

func (m *ProjectModel) Get(id int) (Project, error) {
	var p Project
	err := m.DB.QueryRow("SELECT id, name, created_at FROM projects WHERE id = $1", id).Scan(&p.ID, &p.Name, &p.CreatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return Project{}, ErrNoRecord
	}
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

func (m *ProjectModel) Delete(id int) error {
	res, err := m.DB.Exec("DELETE FROM projects WHERE id = $1", id)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return ErrNoRecord
	}
	return nil
}
