package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Environments struct {
	API_KEY_GEMINI                 string
	GOOGLE_APPLICATION_CREDENTIALS string
	PROJECT_ID                     string
}

func CargarEnv() (Environments, error) {
	environ := Environments{}
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error cargando .env: %v", err)
	}

	environ.API_KEY_GEMINI = os.Getenv("API_KEY_GEMINI")
	environ.GOOGLE_APPLICATION_CREDENTIALS = os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	environ.PROJECT_ID = os.Getenv("PROJECT_ID")

	return environ, nil
}
