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

func (s *TodoTaskService) GetById(listId int) (todo.TodoTask, error) {
	return s.repo.GetById(listId)
}

func (s *TodoTaskService) Delete(listId int) error {
	return s.repo.Delete(listId)
}

func (s *TodoTaskService) Update(listId int, input todo.UpdateTaskInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(listId, input)
}