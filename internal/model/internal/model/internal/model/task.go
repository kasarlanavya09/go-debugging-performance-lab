package model

import "time"

// Task represents a unit of work processed by the service.
type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

// NewTask creates a new task with default values.
func NewTask(id int, title, description string) Task {
	return Task{
		ID:          id,
		Title:       title,
		Description: description,
		Status:      "pending",
		CreatedAt:   time.Now(),
	}
}
