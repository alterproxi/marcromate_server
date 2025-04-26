package config

import (
	"os"

	"github.com/joho/godotenv"
)

func OpenaiApiKey() string {
	return os.Getenv("OPENAI_APIKEY")
}

func LoadEnv() error {
	err := godotenv.Load(".env")
	return err
}
