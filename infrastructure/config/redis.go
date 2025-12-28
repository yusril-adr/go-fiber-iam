package config

import utils "iam-service/infrastructure/utils"

var (
	REDIS_HOST     = utils.GetEnv("REDIS_HOST", "localhost")
	REDIS_PORT     = utils.GetEnv("REDIS_PORT", "6379")
	REDIS_PASSWORD = utils.GetEnv("REDIS_PASSWORD", "")
	REDIS_DB       = utils.GetEnv("REDIS_DB", "0")
)
