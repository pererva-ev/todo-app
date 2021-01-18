package todo

// Column ...
type Column struct {
	ID     uint16 `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}