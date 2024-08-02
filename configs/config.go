package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DBString string
)

func LoadConfig() {
	DBString = getEnv("DB_STRING")
}

// use godot package to load/read the .env file and
// return the value of the key
func getEnv(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf(err.Error())
	}

	return os.Getenv(key)
}
