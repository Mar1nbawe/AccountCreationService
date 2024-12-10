package envFuncs

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var envPrimer = false

func initEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(".env file not found")
	}
	log.Print(".env found. Project continuing.")
	envPrimer = true
}

func GetEnvVar(envvar string) (string, error) {
	if envPrimer == false {
		initEnv()
	}
	value, exists := os.LookupEnv(envvar)
	if !exists {
		return "", fmt.Errorf("environment variable %s not found", envvar)
	}
	return value, nil
}
