package service

import (
	"errors"
	"sync"

	"github.com/kasarlanavya09/go-debugging-performance-lab/internal/model"
)

// TaskService manages tasks safely across concurrent requests.
type TaskService struct {
	mu     sync.RWMutex
	tasks  map[int]model.Task
	nextID int
}

// NewTaskService initializes an empty task service.
func NewTaskService() *TaskService {
	return &TaskService{
		tasks:  make(map[int]model.Task),
		nextID: 1,
	}
}

// CreateTask creates and stores a new task.
func (s *TaskService) CreateTask(title, description string) model.Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	task := model.NewTask(s.nextID, title, description)
	s.tasks[task.ID] = task
	s.nextID++

	return task
}

// GetTask retrieves a task by ID.
func (s *TaskService) GetTask(id int) (model.Task, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	task, exists := s.tasks[id]
	if !exists {
		return model.Task{}, errors.New("task not found")
	}

	return task, nil
}

// GetAllTasks returns all stored tasks.
func (s *TaskService) GetAllTasks() []model.Task {
	s.mu.RLock()
	defer s.mu.RUnlock()

	tasks := make([]model.Task, 0, len(s.tasks))

	for _, task := range s.tasks {
		tasks = append(tasks, task)
	}

	return tasks
}
