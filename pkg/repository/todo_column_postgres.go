package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/pererva-ev/todo-app"
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
	createColumnQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoColumnsTable)
	row := tx.QueryRow(createColumnQuery, column.Name, column.Status)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}


	return id, tx.Commit()
}

func (r *TodoColumnPostgres) GetAll() ([]todo.TodoColumn, error) {
	var columns []todo.TodoColumn

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.column_id ",
		todoColumnsTable)
	err := r.db.Select(&columns, query)

	return columns, err
}

func (r *TodoColumnPostgres) GetById(columnId int) (todo.TodoColumn, error) {
	var column todo.TodoColumn

	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description FROM %s tl
								INNER JOIN %s ul on tl.id = ul.column_id AND ul.column_id = $2`,
		todoColumnsTable)
	err := r.db.Get(&column, query, columnId)

	return column, err
}

func (r *TodoColumnPostgres) Delete(columnId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.column_id AND ul.column_id=$2",
		todoColumnsTable)
	_, err := r.db.Exec(query, columnId)

	return err
}

func (r *TodoColumnPostgres) Update(columnId int, input todo.UpdateColumnInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	if input.Status != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Status)
		argId++
	}

	// title=$1
	// description=$1
	// title=$1, description=$2
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.column_id AND ul.column_id=$%d",
		todoColumnsTable, setQuery, argId, argId+1)
	args = append(args, columnId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
