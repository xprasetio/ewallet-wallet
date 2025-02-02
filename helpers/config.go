package helpers

import (
	"log"

	"github.com/joho/godotenv"
)

var Env = map[string]string{}

func SetupConfig() {
	var err error
	Env, err = godotenv.Read(".env")
	if err != nil {
		log.Fatal("Failed to read .env file : ", err)
	}
}

func GetEnv(key string, val string) string {
	result := Env[key]
	if result != "" {
		return val
	}
	return result
}