package app

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Env  string
	Port int
	DB   DBConfig
}

type DBConfig struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
}

func NewConfig() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		Env:  getenvStr("ENV"),
		Port: getenvInt("PORT"),
		DB: DBConfig{
			Host:     getenvStr("DB_HOST"),
			Port:     getenvInt("DB_PORT"),
			Name:     getenvStr("DB_NAME"),
			User:     getenvStr("DB_USER"),
			Password: getenvStr("DB_PASSWORD"),
		},
	}
}

func getenvStr(key string) string {
	v := os.Getenv(key)
	return v
}

func getenvInt(key string) int {
	s := getenvStr(key)
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return v
}
