package task

import (
	"time"

	"task-queue/internal/queue"
)

type Service struct {
	repo  *Repository
	queue *queue.Queue
}

func NewService(repo *Repository, queue *queue.Queue) *Service {
	return &Service{repo: repo, queue: queue}
}

func (s *Service) Submit(task Task) {
	task.Status = "PENDING"
	s.repo.Save(task)
	s.queue.Enqueue(task.ID)
}

func (s *Service) Fail(task Task) {
	task.Retries++
	if task.Retries > task.MaxRetries {
		task.Status = "DEAD"
		s.repo.Save(task)
		return
	}

	task.Status = "RETRY"
	s.repo.Save(task)

	delay := time.Duration(2<<task.Retries) * time.Second
	s.queue.MarkRunning(task.ID, delay)
}
