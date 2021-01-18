package todo
// Task ...
type Task struct {
	ID              uint16 `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
}