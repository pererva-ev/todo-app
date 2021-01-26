package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/pererva-ev/todo-app"
)

type TodoTaskPostgres struct {
	db *sqlx.DB
}

func NewTodoTaskPostgres(db *sqlx.DB) *TodoTaskPostgres {
	return &TodoTaskPostgres{db: db}
}

func (r *TodoTaskPostgres) Create(task todo.TodoTask) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createTaskQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoTasksTable)
	row := tx.QueryRow(createTaskQuery, task.Name, task.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}


	return id, tx.Commit()
}

func (r *TodoTaskPostgres) GetAll() ([]todo.TodoTask, error) {
	var tasks []todo.TodoTask

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.task_id ",
		todoTasksTable)
	err := r.db.Select(&tasks, query)

	return tasks, err
}

func (r *TodoTaskPostgres) GetById(taskId int) (todo.TodoTask, error) {
	var task todo.TodoTask

	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description FROM %s tl
								INNER JOIN %s ul on tl.id = ul.task_id AND ul.task_id = $2`,
		todoTasksTable)
	err := r.db.Get(&task, query, taskId)

	return task, err
}

func (r *TodoTaskPostgres) Delete(taskId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.task_id AND ul.task_id=$2",
		todoTasksTable)
	_, err := r.db.Exec(query, taskId)

	return err
}

func (r *TodoTaskPostgres) Update(taskId int, input todo.UpdateTaskInput) error {
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

	// title=$1
	// description=$1
	// title=$1, description=$2
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.task_id AND ul.task_id=$%d",
		todoTasksTable, setQuery, argId, argId+1)
	args = append(args, taskId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
