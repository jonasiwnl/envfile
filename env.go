package main

import (
	"errors"

	"github.com/joho/godotenv"
)

type Environment struct {
	DatabaseUrl  string
	DatabaseName string
}

func ValidateEnv() (*Environment, error) {
	// Parse .env
	var env map[string]string
	env, err := godotenv.Read()
	if err != nil {
		return nil, err
	}

	DatabaseUrl := env["DATABASE_URL"]
	DatabaseName := env["DATABASE_NAME"]

	if DatabaseUrl == "" || DatabaseName == "" {
		return nil, errors.New("failed to parse local env")
	}

	return &Environment{DatabaseUrl, DatabaseName}, nil
}
