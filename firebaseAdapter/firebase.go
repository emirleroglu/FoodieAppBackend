package firebaseAdapter

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type firebaseConfig struct {
	Client  *firestore.Client
	Context context.Context
}

func FirebaseInit() firebaseConfig {
	ctx := context.Background()
	sa := option.WithCredentialsFile("/Users/halilemirleroglu/go/Workspace/FoodieAppBackend/deposerver-94a87-faa334baf27f.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	config := firebaseConfig{Client: client,
		Context: ctx}

	return config
}
