package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func InitEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

}

func GetEnv(key string, defaultValue string) string {
	InitEnv()
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
