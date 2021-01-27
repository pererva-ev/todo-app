package service

import (
	"github.com/pererva-ev/todo-app"
	"github.com/pererva-ev/todo-app/pkg/repository"
)


type TodoColumnService struct {
	repo repository.TodoColumn
}

func NewTodoColumnService(repo repository.TodoColumn) *TodoColumnService {
	return &TodoColumnService{repo: repo}
}

func (s *TodoColumnService) Create(list todo.TodoColumn) (int, error) {
	return s.repo.Create(list)
}

func (s *TodoColumnService) GetAll() ([]todo.TodoColumn, error) {
	return s.repo.GetAll()
}

func (s *TodoColumnService) GetById(columnId int) (todo.TodoColumn, error) {
	return s.repo.GetById(columnId)
}

func (s *TodoColumnService) Delete(columnId int) error {
	return s.repo.Delete(columnId)
}

func (s *TodoColumnService) Update(columnId int, input todo.UpdateColumnInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(columnId, input)
}