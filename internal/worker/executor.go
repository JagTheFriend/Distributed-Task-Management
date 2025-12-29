package worker

import (
	"errors"
	"log"
	"time"

	"task-queue/internal/task"
)

func Execute(t task.Task) error {
	log.Println("Executing:", t.ID, t.Payload)
	time.Sleep(time.Second)

	if time.Now().Unix()%2 == 0 {
		return errors.New("simulated failure")
	}
	return nil
}
