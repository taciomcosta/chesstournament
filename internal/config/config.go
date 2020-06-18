package config

import (
	"fmt"
	"github.com/joho/godotenv"
)

var env map[string]string

func init() {
	config, err := godotenv.Read()
	if err != nil {
		fmt.Println("Could not read config from .env")
	}
	env = config
}

func String(key string) string {
	return env[key]
}
