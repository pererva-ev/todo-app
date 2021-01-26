package service

import (
	"github.com/pererva-ev/todo-app"
	"github.com/pererva-ev/todo-app/pkg/repository"
)


type TodoCommentService struct {
	repo repository.TodoComment
}

func NewTodoCommentService(repo repository.TodoComment) *TodoCommentService {
	return &TodoCommentService{repo: repo}
}

func (s *TodoCommentService) Create(list todo.TodoComment) (int, error) {
	return s.repo.Create(list)
}

func (s *TodoCommentService) GetAll() ([]todo.TodoComment, error) {
	return s.repo.GetAll()
}

func (s *TodoCommentService) GetById(listId int) (todo.TodoComment, error) {
	return s.repo.GetById(listId)
}

func (s *TodoCommentService) Delete(listId int) error {
	return s.repo.Delete(listId)
}

func (s *TodoCommentService) Update(listId int, input todo.UpdateCommentInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(listId, input)
}