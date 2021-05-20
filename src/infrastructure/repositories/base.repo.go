package repositories

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
	"log"
)

var CTX = context.Background()

func BaseRepo() *firestore.Client {

	config := &firebase.Config{
		ProjectID: "go-market-26026",
	}

	optionConfig := option.WithCredentialsFile("./config/credentials.json")

	app, err := firebase.NewApp(CTX, config, optionConfig)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Firestore(CTX)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
