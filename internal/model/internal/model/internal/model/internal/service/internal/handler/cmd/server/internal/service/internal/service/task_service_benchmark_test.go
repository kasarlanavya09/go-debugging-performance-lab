package service

import (
	"fmt"
	"testing"
)

func BenchmarkCreateTask(b *testing.B) {
	service := NewTaskService()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		service.CreateTask(
			fmt.Sprintf("Task %d", i),
			"Benchmark task",
		)
	}
}

func BenchmarkGetTask(b *testing.B) {
	service := NewTaskService()

	const taskCount = 1000

	for i := 0; i < taskCount; i++ {
		service.CreateTask(
			fmt.Sprintf("Task %d", i),
			"Benchmark lookup task",
		)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		id := (i % taskCount) + 1

		_, err := service.GetTask(id)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGetAllTasks(b *testing.B) {
	service := NewTaskService()

	for i := 0; i < 1000; i++ {
		service.CreateTask(
			fmt.Sprintf("Task %d", i),
			"Benchmark list task",
		)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		service.GetAllTasks()
	}
}
