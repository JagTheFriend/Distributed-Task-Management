package queue

import (
	"time"

	"task-queue/internal/config"

	"github.com/redis/go-redis/v9"
)

const (
	Pending = "task:pending"
	Running = "task:running"
	Retry   = "task:retry"
	Dead    = "task:dead"
)

type Queue struct {
	rdb *redis.Client
}

func NewQueue(rdb *redis.Client) *Queue {
	return &Queue{rdb: rdb}
}

func (q *Queue) Enqueue(taskID string) {
	q.rdb.LPush(config.Ctx, Pending, taskID)
}

func (q *Queue) Dequeue() string {
	res, _ := q.rdb.BRPop(config.Ctx, 0, Pending).Result()
	return res[1]
}

func (q *Queue) MarkRunning(taskID string, ttl time.Duration) {
	q.rdb.ZAdd(config.Ctx, Running, redis.Z{
		Score:  float64(time.Now().Add(ttl).Unix()),
		Member: taskID,
	})
}

func (q *Queue) RemoveRunning(taskID string) {
	q.rdb.ZRem(config.Ctx, Running, taskID)
}
