package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/kasarlanavya09/go-debugging-performance-lab/internal/service"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

type createTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// HandleTasks handles creating and listing tasks.
func (h *TaskHandler) HandleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.createTask(w, r)
	case http.MethodGet:
		h.getAllTasks(w)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

// HandleTaskByID retrieves an individual task.
func (h *TaskHandler) HandleTaskByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idText := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id, err := strconv.Atoi(idText)

	if err != nil || id <= 0 {
		http.Error(w, "invalid task id", http.StatusBadRequest)
		return
	}

	task, err := h.service.GetTask(id)
	if err != nil {
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}

	writeJSON(w, http.StatusOK, task)
}

func (h *TaskHandler) createTask(w http.ResponseWriter, r *http.Request) {
	var request createTaskRequest

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&request); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	request.Title = strings.TrimSpace(request.Title)
	request.Description = strings.TrimSpace(request.Description)

	if request.Title == "" {
		http.Error(w, "title is required", http.StatusBadRequest)
		return
	}

	task := h.service.CreateTask(request.Title, request.Description)

	writeJSON(w, http.StatusCreated, task)
}

func (h *TaskHandler) getAllTasks(w http.ResponseWriter) {
	tasks := h.service.GetAllTasks()
	writeJSON(w, http.StatusOK, tasks)
}

func writeJSON(w http.ResponseWriter, status int, value interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(value); err != nil {
		return
	}
}
