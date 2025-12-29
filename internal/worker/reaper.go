package worker

import (
	"fmt"
	"time"

	"task-queue/internal/config"
	"task-queue/internal/queue"

	"github.com/redis/go-redis/v9"
)

func StartReaper(rdb *redis.Client) {
	for {
		now := time.Now().Unix()
		expired, _ := rdb.ZRangeByScore(config.Ctx, queue.Running, &redis.ZRangeBy{
			Min: "0",
			Max: fmt.Sprint(now),
		}).Result()

		for _, id := range expired {
			rdb.ZRem(config.Ctx, queue.Running, id)
			rdb.LPush(config.Ctx, queue.Pending, id)
		}
		time.Sleep(3 * time.Second)
	}
}
