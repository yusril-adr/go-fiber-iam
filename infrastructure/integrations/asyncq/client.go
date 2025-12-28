package asyncqintegration

import (
	"github.com/hibiken/asynq"
)

var (
	Client *asynq.Client
)

func InitClient() {
	_, redisClientDbConfig := GetConfig()

	Client = asynq.NewClient(redisClientDbConfig)
}
