package api

import (
	"encoding/json"
	"net/http"

	"task-queue/internal/task"

	"github.com/google/uuid"
)

type Handler struct {
	service *task.Service
}

func NewHandler(service *task.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Type    string `json:"type"`
		Payload string `json:"payload"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	t := task.Task{
		ID:         uuid.NewString(),
		Type:       req.Type,
		Payload:    req.Payload,
		MaxRetries: 3,
	}

	h.service.Submit(t)
	json.NewEncoder(w).Encode(t)
}
