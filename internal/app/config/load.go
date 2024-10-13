package config

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

var (
	config *Config
	once   sync.Once
)

func Load() *Config {
	once.Do(loadConfig)
	return config
}

func loadConfig() {
	err := godotenv.Load()

	log.Println("Loading app config...")

	if err != nil {
		log.Fatal("An error occurred while loading app the config")
	}

	config = &Config{
		Env:       getenvStr("ENV"),
		PublicUrl: getenvStr("PUBLIC_URL"),
		HTTP: HTTPConfig{
			Port: getenvInt("PORT"),
		},
		DB: DBConfig{
			Host:     getenvStr("DB_HOST"),
			Port:     getenvInt("DB_PORT"),
			Name:     getenvStr("DB_NAME"),
			User:     getenvStr("DB_USER"),
			Password: getenvStr("DB_PASSWORD"),
		},
		EmailSender: EmailSenderConfig{
			EmailFrom: getenvStr("EMAIL_SENDER_FROM"),
			Password:  getenvStr("EMAIL_SENDER_PASSWORD"),
		},
	}

	log.Println("App config uploaded successfully")
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
