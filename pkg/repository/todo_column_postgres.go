package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/pererva-ev/todo-app"
	"github.com/sirupsen/logrus"
)

type TodoColumnPostgres struct {
	db *sqlx.DB
}

func NewTodoColumnPostgres(db *sqlx.DB) *TodoColumnPostgres {
	return &TodoColumnPostgres{db: db}
}

func (r *TodoColumnPostgres) Create(column todo.TodoColumn) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createColumnQuery := fmt.Sprintf("INSERT INTO %s (name, status) VALUES ($1, $2) RETURNING id", todoColumnsTable)
	row := tx.QueryRow(createColumnQuery, column.Name, column.Status)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *TodoColumnPostgres) GetAll() ([]todo.TodoColumn, error) {
	var columns []todo.TodoColumn

	query := fmt.Sprintf("SELECT tl.id, tl.name, tl.status FROM %s tl",
		todoColumnsTable)
	err := r.db.Select(&columns, query)

	return columns, err
}

func (r *TodoColumnPostgres) GetById(columnId int) (todo.TodoColumn, error) {
	var column todo.TodoColumn

	query := fmt.Sprintf(`SELECT tl.id, tl.name, tl.status FROM %s tl
							WHERE  tl.comment_id=$%d`,
		todoColumnsTable, columnId)
	err := r.db.Get(&column, query, columnId)

	return column, err
}

func (r *TodoColumnPostgres) Delete(columnId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl  WHERE  tl.comment_id=$%d",
		todoColumnsTable, columnId)
	_, err := r.db.Exec(query, columnId)

	return err
}

func (r *TodoColumnPostgres) Update(columnId int, input todo.UpdateColumnInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	if input.Status != nil {
		setValues = append(setValues, fmt.Sprintf("status=$%d", argId))
		args = append(args, *input.Status)
		argId++
	}

	// name=$1
	// status=$1
	// name=$1, status=$2
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl FROM %s ul WHERE  tl.comment_id=$%d",
		setQuery, todoColumnsTable, argId)
	args = append(args, columnId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
