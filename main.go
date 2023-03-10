package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/pchchv/golog"
)

func init() {
	// Load values from .env into the system
	if err := godotenv.Load(); err != nil {
		golog.Panic("No .env file found")
	}
}

func getEnvValue(v string) string {
	// Getting a value. Outputs a panic if the value is missing
	value, exist := os.LookupEnv(v)
	if !exist {
		golog.Panic("Value %v does not exist", v)
	}
	return value
}

func main() {
	server()
}
