package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/pererva-ev/todo-app"
)

type TodoCommentPostgres struct {
	db *sqlx.DB
}

func NewTodoCommentPostgres(db *sqlx.DB) *TodoCommentPostgres {
	return &TodoCommentPostgres{db: db}
}

func (r *TodoCommentPostgres) Create(comment todo.TodoComment) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createCommentQuery := fmt.Sprintf("INSERT INTO %s (text) VALUES ($1, $2) RETURNING id", todoCommentsTable)
	row := tx.QueryRow(createCommentQuery, comment.Text)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}


	return id, tx.Commit()
}

func (r *TodoCommentPostgres) GetAll() ([]todo.TodoComment, error) {
	var comments []todo.TodoComment

	query := fmt.Sprintf("SELECT tl.id, tl.text FROM %s tl INNER JOIN %s ul on tl.id = ul.comment_id ",
		todoCommentsTable)
	err := r.db.Select(&comments, query)

	return comments, err
}

func (r *TodoCommentPostgres) GetById(commentId int) (todo.TodoComment, error) {
	var comment todo.TodoComment

	query := fmt.Sprintf(`SELECT tl.id, tl.textFROM %s tl
								INNER JOIN %s ul on tl.id = ul.comment_id AND ul.comment_id = $2`,
		todoCommentsTable)
	err := r.db.Get(&comment, query, commentId)

	return comment, err
}

func (r *TodoCommentPostgres) Delete(commentId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.comment_id AND ul.comment_id=$2",
		todoCommentsTable)
	_, err := r.db.Exec(query, commentId)

	return err
}

func (r *TodoCommentPostgres) Update(commentId int, input todo.UpdateCommentInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Text != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Text)
		argId++
	}



	// title=$1
	// description=$1
	// title=$1, description=$2
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.comment_id AND ul.comment_id=$%d",
		todoCommentsTable, setQuery, argId, argId+1)
	args = append(args, commentId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
