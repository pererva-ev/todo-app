package service

import (
	"github.com/pererva-ev/todo-app"
	"github.com/pererva-ev/todo-app/pkg/repository"
	"github.com/pererva-ev/todo-app/todo"
)

type Project interface {
	Create(project todo.TodoProject) (int, error)
	GetAll() ([]todo.TodoProject, error)
	GetById(projectID int) (todo.TodoProject, error)
	Update(projectID int, input todo.UpdateProjectInput) error
	Delete(projectID int) error
}

type Column interface {
	Create(column todo.TodoColumn) (int, error)
	GetAll() ([]todo.TodoColumn, error)
	GetById(columnID int) (todo.TodoColumn, error)
	Update(columnID int, input todo.UpdateColumnInput) error
	Delete(columnID int) error
}

type Task interface {
	Create(project todo.TodoTask) (int, error)
	GetAll() ([]todo.TodoTask, error)
	GetById(taskID int) (todo.TodoTask, error)
	Update(taskID int, input todo.UpdateTaskInput) error
	Delete(taskID int) error
}

type Comment interface {
	Create(project todo.TodoComment) (int, error)
	GetAll() ([]todo.TodoComment, error)
	GetById(commentID int) (todo.TodoComment, error)
	Update(commentID int, input todo.UpdateCommentInput) error
	Delete(commentID int) error
}

type Service struct {
	TodoProject
	TodoColumn
	TodoTask
	TodoComment
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Project: NewProjectService(repos.Service),
		Column:  NewColumnService(repos.Column),
		Task:    NewTaskService(repos.Task),
		Comment: NewCommentService(repos.Comment),
	}
}
