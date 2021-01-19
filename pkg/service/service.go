package service

import (
	"github.com/pererva-ev/todo-app"
	"github.com/pererva-ev/todo-app/pkg/repository"
	"github.com/pererva-ev/todo-app/todo"
)

type Project interface {
	Create(project todo.Project) (int, error)
	GetAll() ([]todo.Project, error)
	GetById(projectID int) (todo.Project, error)
	Update(projectID int, input todo.UpdateProjectInput) error
	Delete(projectID int) error
}

type Column interface {
	Create(column todo.Column) (int, error)
	GetAll() ([]todo.Column, error)
	GetById(columnID int) (todo.Column, error)
	Update(columnID int, input todo.UpdateColumnInput) error
	Delete(columnID int) error
}

type Task interface {
	Create(project todo.Task) (int, error)
	GetAll() ([]todo.Task, error)
	GetById(taskID int) (todo.Task, error)
	Update(taskID int, input todo.UpdateTaskInput) error
	Delete(taskID int) error
}

type Comment interface {
	Create(project todo.Comment) (int, error)
	GetAll() ([]todo.Comment, error)
	GetById(commentID int) (todo.Comment, error)
	Update(commentID int, input todo.UpdateCommentInput) error
	Delete(commentID int) error
}

type Service struct {
	Project
	Column
	Task
	Comment
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Project: NewProjectService(repos.Service),
		Column:  NewColumnService(repos.Column),
		Task:    NewTaskService(repos.Task),
		Comment: NewCommentService(repos.Comment),
	}
}
