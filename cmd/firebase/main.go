package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()

	// 1. Initialize Firebase Admin SDK
	// opt := option.WithCredentialsFile("serviceAccountKey.json")
	opt := option.WithAuthCredentialsFile(option.ServiceAccount, "key.json")

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("Ошибка создания приложения Firebase: %v", err)
	}

	// 2. Get the Messaging client
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	token := "REDACTED"

	// 3. Define the Message payload
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: "Hello from Go!",
			Body:  "This is a push notification sent via FCM.",
		},
		Token: token,
	}

	// 4. Send the message
	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Successfully sent message:", response)
}
