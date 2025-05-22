package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT        string
	DB_PORT     string
	DB_HOST     string
	DB_USER     string
	DB_NAME     string
	DB_PASSWORD string
}

var AppConf Config

func Init() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	AppConf = Config{
		":8080",
		GetEnvOrPanic("DB_PORT"),
		GetEnvOrPanic("DB_HOST"),
		GetEnvOrPanic("DB_USER"),
		GetEnvOrPanic("DB_NAME"),
		GetEnvOrPanic("DB_PASSWORD"),
	}

	return nil
}

func (c Config) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.DB_USER,
		c.DB_PASSWORD,
		c.DB_HOST,
		c.DB_PORT,
		c.DB_NAME,
	)
}

func GetEnvOrPanic(key string) string {
	value := os.Getenv(key)

	if value == "" {
		panic(fmt.Sprintf("❌ Missing required environment variable: %s", key))
	}

	return value
}
