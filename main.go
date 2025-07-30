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
	"bot-restaurante/utils"

	"github.com/google/generative-ai-go/genai"
)

func main() {
	var menuItems []database.Menu
	var geminiClient *genai.Client

	environ, err := utils.CargarEnv()
	if err != nil {
		return
	}

	ctx := context.Background()

	client, err := database.InitializeFirestore(&environ, ctx)
	if err != nil {
		return
	}
	defer client.Close()

	menuItems, err = database.GetAllMenuItems(ctx, client)
	if err != nil {
		log.Fatalf("Error obteniendo items del menú: %v", err)
	}

	fmt.Println("🟢 Menú cargado exitosamente:")
	fmt.Println("asds : ", menuItems)

	geminiClient, err = ia.ConnectIA(ctx, &environ)
	if err != nil {
		log.Fatalf("Error conectando a Gemini: %v", err)
	}
	fmt.Println("🟢 Conexión a Gemini establecida exitosamente")
	model := geminiClient.GenerativeModel("gemini-2.0-flash")

	defer geminiClient.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("🟢 Bot listo. Escribe tu consulta (ej: '¿Cuánto cuesta el pollo a la brasa?') o 'salir':")

	jsonMenu, err := json.Marshal(menuItems)
	if err != nil {
		fmt.Println("Error convirtiendo menuItems a JSON:", err)
		return
	}

	fmt.Println("Menú en formato JSON:", string(jsonMenu))

	for {
		fmt.Print("\n👤 Usuario: ")
		userQuery, _ := reader.ReadString('\n')
		userQuery = strings.TrimSpace(userQuery)

		if strings.ToLower(userQuery) == "salir" {
			break
		}

		answer := ia.GenerateRespuesta(string(jsonMenu), ctx, model, userQuery)
		fmt.Printf("\n🤖 Bot: %s\n", answer)
	}
}
