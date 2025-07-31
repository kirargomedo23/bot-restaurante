package utils

import (
	Interfaces "bot-restaurante/interfaces"
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvironment() (Interfaces.Environments, error) {
	env := Interfaces.Environments{}
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
