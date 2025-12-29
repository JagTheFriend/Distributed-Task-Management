package task

import (
	"encoding/json"

	"task-queue/internal/config"

	"github.com/redis/go-redis/v9"
)

type Repository struct {
	rdb *redis.Client
}

func NewRepository(rdb *redis.Client) *Repository {
	return &Repository{rdb: rdb}
}

func (r *Repository) Save(task Task) {
	data, _ := json.Marshal(task)
	r.rdb.Set(config.Ctx, "task:"+task.ID, data, 0)
}

func (r *Repository) Get(id string) Task {
	val, _ := r.rdb.Get(config.Ctx, "task:"+id).Result()
	var t Task
	json.Unmarshal([]byte(val), &t)
	return t
}
