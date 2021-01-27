package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/pererva-ev/todo-app"
)

type TodoProjectPostgres struct {
	db *sqlx.DB
}

func NewTodoProjectPostgres(db *sqlx.DB) *TodoProjectPostgres {
	return &TodoProjectPostgres{db: db}
}

func (r *TodoProjectPostgres) Create(project todo.TodoProject) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createProjectQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoProjectsTable)
	row := tx.QueryRow(createProjectQuery, project.Name, project.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}


	return id, tx.Commit()
}

func (r *TodoProjectPostgres) GetAll() ([]todo.TodoProject, error) {
	var projects []todo.TodoProject

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl ",
		todoProjectsTable)
	err := r.db.Select(&projects, query)

	return projects, err
}

func (r *TodoProjectPostgres) GetById(projectId int) (todo.TodoProject, error) {
	var project todo.TodoProject

	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description FROM %s tl
								WHERE  tl.project_id=$2`,
		todoProjectsTable)
	err := r.db.Get(&project, query, projectId)

	return project, err
}

func (r *TodoProjectPostgres) Delete(projectId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl  WHERE  tl.project_id=$2",
		todoProjectsTable)
	_, err := r.db.Exec(query, projectId)

	return err
}

func (r *TodoProjectPostgres) Update(projectId int, input todo.UpdateProjectInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl FROM %s ul WHERE tl.id = ul.project_id AND ul.project_id=$%d",
	setQuery, todoProjectsTable, argId)
	args = append(args, projectId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
