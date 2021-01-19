package todo

import "errors"

type TodoColumn struct {
	ID     uint16 `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type TodoComment struct {
	ID   uint16 `json:"id"`
	Text string `json:"text"`
}

type TodoProject struct {
	ID          uint16 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type TodoTask struct {
	ID          uint16 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateColumnInput struct {
	Name   *string `json:"name"`
	Status *string `json:"status"`
}

func (i UpdateColumnInput) Validate() error {
	if i.Name == nil && i.Status == nil {
		return errors.New("update structure has no values")
	}
	return nil
}

type UpdateTaskInput struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

func (i UpdateTaskInput) Validate() error {
	if i.Name == nil && i.Description == nil {
		return errors.New("update structure has no values")
	}
	return nil
}

type UpdateCommentInput struct {
	Name *string `json:"name"`
	Text *string `json:"text"`
}

func (i UpdateCommentInput) Validate() error {
	if i.Name == nil && i.Text == nil {
		return errors.New("update structure has no values")
	}
	return nil
}

type UpdateProjectInput struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

func (i UpdateProjectInput) Validate() error {
	if i.Name == nil && i.Description == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
