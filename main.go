package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"bot-restaurante/database"
	"bot-restaurante/ia"
	Interfaces "bot-restaurante/interfaces"
	"bot-restaurante/utils"

	"github.com/google/generative-ai-go/genai"
)

func main() {
	var menuItems []Interfaces.Menu
	var geminiClient *genai.Client

	environ, err := utils.LoadEnvironment()
	if err != nil {
		return
	}

	ctx := context.Background()

	client, err := database.InitializeFirestore(&environ, ctx)
	if err != nil {
		return
	}
	defer client.Close()

	menuItems, err = database.GetAllMenuActive(ctx, client)
	if err != nil {
		log.Fatalf("Error obteniendo items del menú: %v", err)
		return
	}

	fmt.Println("🟢 Menú cargado exitosamente:")

	geminiClient, err = ia.ConnectIA(ctx, &environ)
	if err != nil {
		log.Fatalf("Error conectando a Gemini: %v", err)
		return
	}
	fmt.Println("🟢 Conexión a Gemini establecida exitosamente")
	model := geminiClient.GenerativeModel("gemini-2.0-flash")

	defer geminiClient.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("🟢 Bot listo ✅✅✅")

	jsonMenu, err := json.Marshal(menuItems)
	if err != nil {
		fmt.Println("Error convirtiendo menuItems a JSON:", err)
		return
	}

	for {
		fmt.Print("\n👤 Usuario: ")
		userQuery, _ := reader.ReadString('\n')
		userQuery = strings.TrimSpace(userQuery)

		if strings.ToLower(userQuery) == "esc" {
			break
		}

		typeQuestion, err := ia.CategorizeQuestion(ctx, model, userQuery)
		if err != nil {
			fmt.Println("Error al categorizar la pregunta:", err)
			return
		}
		fmt.Printf("🤖🤖🤖 Bot: \n Tipo de pregunta: %s\n", typeQuestion)
		answer, err := ia.GenerateAnswer(typeQuestion, string(jsonMenu), ctx, model, userQuery)
		if err != nil {
			fmt.Println("Error al generar respuestas:", err)
			return
		}
		fmt.Println("🤖🤖🤖 Bot: ")
		fmt.Printf("%s", answer)

	}
}
