package todo

// Project ...
type Project struct {
	ID          uint16 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}