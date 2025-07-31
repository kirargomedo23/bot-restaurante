package database

import (
	"context"
	"fmt"

	Interfaces "bot-restaurante/interfaces"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func InitializeFirestore(environment *Interfaces.Environments, ctx context.Context) (*firestore.Client, error) {
	conf := &firebase.Config{ProjectID: environment.PROJECT_ID}
	opt := option.WithCredentialsFile(environment.GOOGLE_APPLICATION_CREDENTIALS)

	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		return nil, fmt.Errorf("Error initializing Firebase: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, fmt.Errorf("Error connecting to Firestore: %v", err)
	}

	return client, nil
}

func GetAllMenuActive(ctx context.Context, client *firestore.Client) ([]Interfaces.Menu, error) {
	doc, err := client.Collection("restaurant_mock").Doc("bhHn47g5CQV0hyt7GPYc").Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("Error obteniendo documento 'menu': %v", err)
	}
	data := doc.Data()
	var menu []Interfaces.Menu
	for _, v := range data {
		m, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		active, ok := m["active"].(bool)
		if !ok || !active {
			continue
		}

		plato := Interfaces.Menu{
			Name:        fmt.Sprintf("%v", m["name"]),
			Price:       m["price"].(float64),
			Description: fmt.Sprintf("%v", m["description"]),
			Active:      active,
		}
		menu = append(menu, plato)
	}

	return menu, nil
}
