package ia

import (
	"context"
	"fmt"
	"log"

	"bot-restaurante/utils"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func ConnectIA(ctx context.Context, environ *utils.Environments) (*genai.Client, error) {
	genIA, err := genai.NewClient(ctx, option.WithAPIKey(environ.API_KEY_GEMINI))
	if err != nil {
		log.Fatalf("Error al conectar con Gemini: %v", err)
	}

	return genIA, nil
}

func GenerateRespuesta(menu string, ctx context.Context, model *genai.GenerativeModel, query string) string {
	prompt := fmt.Sprintf(`
		Eres un asistente de restaurante. Basado en este menú: %+v, 
		responde esta consulta: "%s". 
		Responde máximo en 2 líneas con formato:
		"[PLATO]: S/[PRECIO] - [DESCRIPCIÓN]"
	`, menu, query)

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Printf("Error en Gemini: %v", err)
		return "Lo siento, ocurrió un error al procesar tu consulta"
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return "No encontré información sobre ese plato"
	}

	return fmt.Sprintf("%v", resp.Candidates[0].Content.Parts[0])
}
