package utils

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Environments struct {
	API_KEY_GEMINI                 string
	GOOGLE_APPLICATION_CREDENTIALS string
	PROJECT_ID                     string
}

func LoadEnvironment() (Environments, error) {
	env := Environments{}
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error cargando .env: %v", err)
		return env, errors.New("Error cargando .env : " + err.Error())
	}

	env.API_KEY_GEMINI = os.Getenv("API_KEY_GEMINI")
	env.GOOGLE_APPLICATION_CREDENTIALS = os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	env.PROJECT_ID = os.Getenv("PROJECT_ID")

	return env, nil
}
