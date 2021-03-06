package service

import (
	"github.com/pererva-ev/todo-app"
	"github.com/pererva-ev/todo-app/pkg/repository"
)


type TodoProjectService struct {
	repo repository.TodoProject
}

func NewTodoProjectService(repo repository.TodoProject) *TodoProjectService {
	return &TodoProjectService{repo: repo}
}

func (s *TodoProjectService) Create(list todo.TodoProject) (int, error) {
	return s.repo.Create(list)
}

func (s *TodoProjectService) GetAll() ([]todo.TodoProject, error) {
	return s.repo.GetAll()
}

func (s *TodoProjectService) GetById(projectId int) (todo.TodoProject, error) {
	return s.repo.GetById(projectId)
}

func (s *TodoProjectService) Delete(projectId int) error {
	return s.repo.Delete(projectId)
}

func (s *TodoProjectService) Update(projectId int, input todo.UpdateProjectInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(projectId, input)
}