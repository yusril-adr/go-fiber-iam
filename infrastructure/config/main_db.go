package config

import utils "iam-service/infrastructure/utils"

var (
	MAIN_DB_HOST                          = utils.GetEnv("MAIN_DB_HOST", "")
	MAIN_DB_PORT                          = utils.GetEnv("MAIN_DB_PORT", "")
	MAIN_DB_USER                          = utils.GetEnv("MAIN_DB_USER", "")
	MAIN_DB_PASS                          = utils.GetEnv("MAIN_DB_PASS", "")
	MAIN_DB_SSL_MODE                      = utils.GetEnv("MAIN_DB_SSL_MODE", "false")
	MAIN_DB_NAME                          = utils.GetEnv("MAIN_DB_NAME", "")
	MAIN_DB_TIMEZONE                      = utils.GetEnv("MAIN_DB_TIMEZONE", "UTC")
	MAIN_DB_MAX_IDLE_CONNS                = utils.GetEnv("MAIN_DB_MAX_IDLE_CONNS", "5")
	MAIN_DB_MAX_OPEN_CONNS                = utils.GetEnv("MAIN_DB_MAX_OPEN_CONNS", "10")
	MAIN_DB_MAX_IDLE_CONNS_IN_MINUTES     = utils.GetEnv("MAIN_DB_MAX_IDLE_CONNS_IN_MINUTES", "10")
	MAIN_DB_MAX_LIFETIME_CONNS_IN_MINUTES = utils.GetEnv("MAIN_DB_MAX_LIFETIME_CONNS_IN_MINUTES", "60")
)
