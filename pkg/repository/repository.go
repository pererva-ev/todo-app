package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/pererva-ev/todo-app"
)

type TodoProject interface {
	Create(project todo.TodoProject) (int, error)
	GetAll([]todo.TodoProject, error)
	GetById(projectID int) (todo.TodoProject, error)
	Update(projectID int, input todo.UpdateProjectInput) error
	Delete(projectID int) error
}

type TodoColumn interface {
	Create(column todo.TodoColumn) (int, error)
	GetAll([]todo.TodoColumn, error)
	GetById(columnID int) (todo.TodoColumn, error)
	Update(columnID int, input todo.UpdateColumnInput) error
	Delete(columnID int) error
}

type TodoTask interface {
	Create(project todo.TodoTask) (int, error)
	GetAll([]todo.TodoTask, error)
	GetById(taskID int) (todo.TodoTask, error)
	Update(taskID int, input todo.UpdateTaskInput) error
	Delete(taskID int) error
}

type TodoComment interface {
	Create(project todo.TodoComment) (int, error)
	GetAll([]todo.TodoComment, error)
	GetById(commentID int) (todo.TodoComment, error)
	Update(commentID int, input todo.UpdateCommentInput) error
	Delete(commentID int) error
}

type Repository struct {
	TodoProject
	TodoColumn
	TodoTask
	TodoComment
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		TodoProject: NewTodoProjectPosgres(db),
		TodoColumn:  NewTodoColumnPosgres(db),
		TodoTask:    NewTodoTaskPostgres(db),
		TodoComment: NewTodoCommentPosgres(db),
	}
}
