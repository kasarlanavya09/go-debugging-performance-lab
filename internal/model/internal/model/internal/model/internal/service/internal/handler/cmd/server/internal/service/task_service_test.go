package service

import (
	"sync"
	"testing"
)

func TestCreateTask(t *testing.T) {
	service := NewTaskService()

	task := service.CreateTask("Fix API bug", "Investigate timeout issue")

	if task.ID != 1 {
		t.Fatalf("expected task ID 1, got %d", task.ID)
	}

	if task.Title != "Fix API bug" {
		t.Errorf("expected title 'Fix API bug', got %q", task.Title)
	}

	if task.Status != "pending" {
		t.Errorf("expected status 'pending', got %q", task.Status)
	}
}

func TestGetTask(t *testing.T) {
	service := NewTaskService()

	created := service.CreateTask("Refactor service", "Improve maintainability")

	task, err := service.GetTask(created.ID)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if task.ID != created.ID {
		t.Errorf("expected task ID %d, got %d", created.ID, task.ID)
	}
}

func TestGetTaskNotFound(t *testing.T) {
	service := NewTaskService()

	_, err := service.GetTask(999)

	if err == nil {
		t.Fatal("expected error for missing task")
	}
}

func TestConcurrentTaskCreation(t *testing.T) {
	service := NewTaskService()

	const workers = 100

	var wg sync.WaitGroup
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			service.CreateTask("Concurrent task", "Created by goroutine")
		}()
	}

	wg.Wait()

	tasks := service.GetAllTasks()

	if len(tasks) != workers {
		t.Fatalf("expected %d tasks, got %d", workers, len(tasks))
	}

	ids := make(map[int]bool)

	for _, task := range tasks {
		if ids[task.ID] {
			t.Fatalf("duplicate task ID detected: %d", task.ID)
		}

		ids[task.ID] = true
	}
}
