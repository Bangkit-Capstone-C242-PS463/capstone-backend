package utils

import (
	"github.com/joho/godotenv"
	"os"
)

func DetermineEnvName(env string) string {
	switch env {
	case "prod":
		return "prod"
	default:
		return "dev"
	}
}

func LoadEnvFile(fileName string) bool {
	if err := godotenv.Load(fileName, ".env.dev"); err != nil {
		return false
	}
	return true
}

func IS_PROD() bool  { return os.Getenv("ENV") == "production" }
func IS_DEV() bool   { return os.Getenv("ENV") == "development" }
