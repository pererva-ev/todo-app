package service

import (
	"github.com/pererva-ev/todo-app"
	"github.com/pererva-ev/todo-app/pkg/repository"
)


type TodoTaskService struct {
	repo repository.TodoTask
}

func NewTodoTaskService(repo repository.TodoTask) *TodoTaskService {
	return &TodoTaskService{repo: repo}
}

func (s *TodoTaskService) Create(list todo.TodoTask) (int, error) {
	return s.repo.Create(list)
}

func (s *TodoTaskService) GetAll() ([]todo.TodoTask, error) {
	return s.repo.GetAll()
}

func (s *TodoTaskService) GetById(taskId int) (todo.TodoTask, error) {
	return s.repo.GetById(taskId)
}

func (s *TodoTaskService) Delete(taskId int) error {
	return s.repo.Delete(taskId)
}

func (s *TodoTaskService) Update(taskId int, input todo.UpdateTaskInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(taskId, input)
}