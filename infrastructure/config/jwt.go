package config

import "iam-service/infrastructure/utils"

var (
	JWT_ACCESS_TOKEN_SECRET      = utils.GetEnv("JWT_ACCESS_TOKEN_SECRET", "secret")
	JWT_ACCESS_TOKEN_EXPIRES_IN  = utils.GetEnv("JWT_ACCESS_TOKEN_EXPIRES_IN", "30m")
	JWT_REFRESH_TOKEN_SECRET     = utils.GetEnv("JWT_REFRESH_TOKEN_SECRET", "secret")
	JWT_REFRESH_TOKEN_EXPIRES_IN = utils.GetEnv("JWT_REFRESH_TOKEN_EXPIRES_IN", "24h")
)
