package main

import (
	"log"
	"time"

	"task-queue/internal/config"
	"task-queue/internal/queue"
	"task-queue/internal/task"
	"task-queue/internal/worker"
)

func main() {
	rdb := config.NewRedis()
	repo := task.NewRepository(rdb)
	queue := queue.NewQueue(rdb)

	go worker.StartReaper(rdb)

	log.Println("Worker started")
	for {
		taskID := queue.Dequeue()
		t := repo.Get(taskID)

		queue.MarkRunning(taskID, 10*time.Second)
		err := worker.Execute(t)
		queue.RemoveRunning(taskID)

		if err != nil {
			t.Retries++
			repo.Save(t)
			queue.Enqueue(taskID)
		} else {
			t.Status = "DONE"
			repo.Save(t)
		}
	}
}
