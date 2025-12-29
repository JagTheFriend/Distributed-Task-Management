package main

import (
	"log"
	"net/http"

	"task-queue/internal/api"
	"task-queue/internal/config"
	"task-queue/internal/queue"
	"task-queue/internal/task"
)

func main() {
	rdb := config.NewRedis()
	repo := task.NewRepository(rdb)
	queue := queue.NewQueue(rdb)
	service := task.NewService(repo, queue)
	handler := api.NewHandler(service)

	http.HandleFunc("/tasks", handler.CreateTask)
	log.Println("API running on :8080")
	http.ListenAndServe(":8080", nil)
}
