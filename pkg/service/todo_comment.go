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

func (s *TodoCommentService) GetById(commentId int) (todo.TodoComment, error) {
	return s.repo.GetById(commentId)
}

func (s *TodoCommentService) Delete(commentId int) error {
	return s.repo.Delete(commentId)
}

func (s *TodoCommentService) Update(commentId int, input todo.UpdateCommentInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(commentId, input)
}