package main

import (
	"log"
	"net/http"

	"github.com/kasarlanavya09/go-debugging-performance-lab/internal/handler"
	"github.com/kasarlanavya09/go-debugging-performance-lab/internal/service"
)

func main() {
	taskService := service.NewTaskService()
	taskHandler := handler.NewTaskHandler(taskService)

	mux := http.NewServeMux()

	mux.HandleFunc("/tasks", taskHandler.HandleTasks)
	mux.HandleFunc("/tasks/", taskHandler.HandleTaskByID)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Task API running on http://localhost:8080")

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server failed: %v", err)
	}
}
