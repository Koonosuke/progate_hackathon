package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser string
	DBPassword string
	DBHost string
	DBPort string
	DBName string
	OpenAIApiKey string
	ServerPort string
	MiroAPIToken string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println(".envファイルが読み込めませんでした。")
	}

	return &Config{
        DBUser:       os.Getenv("DB_USER"),
        DBPassword:   os.Getenv("DB_PASSWORD"),
        DBHost:       os.Getenv("DB_HOST"),
        DBPort:       os.Getenv("DB_PORT"),
        DBName:       os.Getenv("DB_NAME"),
        OpenAIApiKey: os.Getenv("OPENAIAPI_KEY"),
        ServerPort:   os.Getenv("SERVER_PORT"),
		MiroAPIToken: os.Getenv("MIRO_ACCESS_TOKEN"),
    }
}