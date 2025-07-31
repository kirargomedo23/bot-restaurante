package ia

import (
	"context"
	"fmt"
	"log"

	Interfaces "bot-restaurante/src/interfaces"
	"errors"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func ConnectIA(ctx context.Context, environ *Interfaces.Environments) (*genai.Client, error) {
	genIA, err := genai.NewClient(ctx, option.WithAPIKey(environ.API_KEY_GEMINI))
	if err != nil {
		log.Fatalf("Error al conectar con Gemini: %v", err)
		return nil, errors.New("Error al conectar con Gemini : " + err.Error())
	}

	return genIA, nil
}

func CategorizeQuestion(ctx context.Context, model *genai.GenerativeModel, question string) (string, error) {

	prompt := fmt.Sprintf(`
		Eres un clasificador experto en consultas de restaurante. Analiza la siguiente pregunta y clasifícala SOLO con una de estas etiquetas:

		1. "INFO_PLATO" - Cuando:
		- Pregunta por la existencia de un plato (ej: "¿tienen lasagna?", "hay menu vegetariano?")
		- Solicita detalles (ej: "¿el ceviche lleva apio?", "el lomo es picante?")
		- Usa verbos como: tener, haber, incluir, llevar, contener

		2. "CALCULO_PRECIO" - Cuando:
		- Menciona cantidades numéricas (ej: "2 pollos", "quiero 3")
		- Usa palabras como: total, costo, cuánto por, valor, precio de X unidades
		- Contiene operaciones (ej: "para 4 personas", "mitad y mitad")

		3. "OTRO" - Para:
		- Horarios, ubicación, reservas
		- Promociones no especificadas
		- Consultas no relacionadas al menú

		Ejemplos:
		- "¿Tienen pastel de chocolate?" → "INFO_PLATO"
		- "Dos ceviches y una chicha" → "CALCULO_PRECIO"
		- "A qué hora cierran?" → "OTRO"

		Consulta a clasificar: "%s"

		RESPONDE EXCLUSIVAMENTE CON LA ETIQUETA ENTRE COMILLAS (ej: "CALCULO_PRECIO")
		`, question)

	answer, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatalf("Error al obtener la clasificacion con Gemini: %v", err)
		return "", errors.New("Error al obtener la clasificacion con Gemini : " + err.Error())
	}

	response := answer.Candidates[0].Content.Parts[0]
	str := fmt.Sprintf("%v", response)
	str = strings.TrimSpace(str)
	str = strings.ReplaceAll(str, "\"", "")
	return strings.TrimSpace(fmt.Sprintf("%v", str)), nil
}

func GenerateAnswer(typeQuestion string, menu string, ctx context.Context, model *genai.GenerativeModel, query string) (string, error) {

	switch typeQuestion {
	case "INFO_PLATO":
		return GenerateAnswerWithMenu(menu, ctx, model, query)
	case "CALCULO_PRECIO":
		return calculatePrice(menu, query, ctx, model)
	default:
		return "¿Podrías especificar si quieres información de un plato o calcular un precio?", nil
	}

}

func GenerateAnswerWithMenu(menu string, ctx context.Context, model *genai.GenerativeModel, query string) (string, error) {
	prompt := fmt.Sprintf(`
		Eres un asistente de restaurante. Basado en este menú: %+v, 
		responde esta consulta: "%s". 
		Responde máximo en 2 líneas con formato:
		"- Menu🍽️ : [PLATO]"
		"- Precio 💵 : S/[PRECIO]  "
		"- Descripcion: [DESCRIPCIÓN] "
	`, menu, query)

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Printf("Error en Gemini: %v", err)
		return "", errors.New("Error con Gemini : " + err.Error())
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return "No encontré información sobre ese plato", nil
	}

	return fmt.Sprintf("%v", resp.Candidates[0].Content.Parts[0]), nil

}

func calculatePrice(menu string, query string, ctx context.Context, model *genai.GenerativeModel) (string, error) {
	prompt := fmt.Sprintf(`
		Eres un asistente de restaurante. Basado en este menú: %+v, 
		responde esta consulta: "%s". 
		Responde máximo en 2 líneas con formato:
		"- [PLATO] x [CANTIDAD]: S/[TOTAL] " 
	`, menu, query)

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Printf("Error en Gemini: %v", err)
		return "", errors.New("Error con Gemini : " + err.Error())
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return "No encontré información sobre ese plato", nil
	}

	return fmt.Sprintf("%v", resp.Candidates[0].Content.Parts[0]), nil
}
