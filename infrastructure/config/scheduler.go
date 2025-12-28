package config

import (
	"iam-service/infrastructure/utils"
)

var (
	SCHEDULER_CONCURRENCY = utils.GetEnv("SCHEDULER_CONCURRENCY", "10")
)
