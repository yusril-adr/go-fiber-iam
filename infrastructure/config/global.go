package config

import (
	"iam-service/constants"
	"iam-service/infrastructure/utils"
)

var (
	PORT    = utils.GetEnv("PORT", "3000")
	APP_ENV = utils.GetEnv("APP_ENV", constants.APP_ENV_DEV)
)
