package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	API_HOST string
	API_PORT string
}

var ENV Config

func initConfig() {
	godotenv.Load()

	ENV.API_HOST = os.Getenv("API_HOST")
	ENV.API_PORT = os.Getenv("API_PORT")
}