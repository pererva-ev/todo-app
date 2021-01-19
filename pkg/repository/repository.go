package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/pererva-ev/todo-app"
)

type Project interface {
	Create(project todo.Project) (int, error)
	GetAll([]todo.Project, error)
	GetById(projectID int) (todo.Project, error)
	Update(projectID int, input todo.UpdateProjectInput) error
	Delete(projectID int) error
}

type Column interface {
	Create(column todo.Column) (int, error)
	GetAll([]todo.Column, error)
	GetById(columnID int) (todo.Column, error)
	Update(columnID int, input todo.UpdateColumnInput) error
	Delete(columnID int) error
}

type Task interface {
	Create(project todo.Task) (int, error)
	GetAll([]todo.Task, error)
	GetById(taskID int) (todo.Task, error)
	Update(taskID int, input todo.UpdateTaskInput) error
	Delete(taskID int) error
}

type Comment interface {
	Create(project todo.Comment) (int, error)
	GetAll([]todo.Comment, error)
	GetById(commentID int) (todo.Comment, error)
	Update(commentIDprojectId int, input todo.UpdateCommentInput) error
	Delete(commentID int) error
}

type Repository interface {
	Project
	Column
	Task
	Comment
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Project: NewProjectPosgres(db),
		Column:  NewColumnPosgres(db),
		Task:    NewTaskPosgres(db),
		Comment: NewCmmentPosgres(db),
	}
}
