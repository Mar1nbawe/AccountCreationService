package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func initEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(".env file not found")
	}
	log.Print(".env found. Project continuing.")
}

func getEnvVar(envvar string) (string, error) {
	initEnv()
	value, exists := os.LookupEnv(envvar)
	if !exists {
		return "", fmt.Errorf("environment variable %s not found", envvar)
	}
	return value, nil
}
