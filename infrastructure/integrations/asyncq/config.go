package asyncqintegration

import (
	"fmt"
	"strconv"

	"github.com/hibiken/asynq"

	"iam-service/infrastructure/config"
)

func GetConfig() (asynq.Config, asynq.RedisClientOpt) {
	redisClientDbConfig, err := strconv.Atoi(config.REDIS_DB)
	if err != nil {
		panic(err.Error())
	}
	schedulerConcurrentConfig, err := strconv.Atoi(config.SCHEDULER_CONCURRENCY)
	if err != nil {
		panic(err.Error())
	}

	asyncqRedisConfig := asynq.RedisClientOpt{
		Addr: fmt.Sprintf("%s:%s", config.REDIS_HOST, config.REDIS_PORT),
		// Omit if no password is required
		Password: config.REDIS_PASSWORD,
		// Use a dedicated db number for asynq.
		// By default, Redis offers 16 databases (0..15)
		DB: redisClientDbConfig,
	}

	var asynqConfig = asynq.Config{
		// Specify how many concurrent workers to use
		Concurrency: schedulerConcurrentConfig,
		// Optionally specify multiple queues with different priority.
		Queues: map[string]int{
			"critical": 6,
			"default":  3,
			"low":      1,
		},
		// See the godoc for other configuration options
	}

	return asynqConfig, asyncqRedisConfig
}
