package main

import (
	"fmt"
	"log"
	"strconv"

	"emirleroglu.com/foodie/firebaseAdapter"
	"emirleroglu.com/foodie/model"
)

func main() {

	firedge := model.NewFridge("TestDeposu", 12)

	config := firebaseAdapter.FirebaseInit()
	//firedge := fridge.NewFridge("Halil", 2)
	//fmt.Println(firedge.Name)
	//fmt.Println(firedge.ID)

	//client.Collection("fridges").Add(ctx, firedge)
	s2 := strconv.Itoa(firedge.ID)
	config.Client.Collection("fridges").Doc(s2).Set(config.Context, firedge)

	query := config.Client.Collection("fridges").Where("ID", "==", 2).Documents(config.Context)

	doc, err := query.Next()
	if err != nil {
		log.Fatal("Failed to check")
	}
	fmt.Println(doc.Data())
	myFridge := model.AddIngredientsToFrifgeId(2)
	config.Client.Collection("fridges").Doc("3").Set(config.Context, myFridge)

	/*client.Collection("users").Add(ctx, map[string]interface{}{
		"first": "Ada",
		"last":  "Lovelace",
		"born":  1815,
	})
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}

	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}
	if err != nil {
		log.Fatalln(err)
	}

	// ...

	client.Collection("users").Doc("frank").Set(ctx, map[string]interface{}{
		"name": "Frank",
		"age":  13,
		"favorites": map[string]interface{}{
			"food":    "Pizza",
			"color":   "Blue",
			"subject": "recess",
		},
	}, firestore.MergeAll)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}

	/*	iter := client.Collection("users").Documents(ctx)
		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				log.Fatalf("Failed to iterate: %v", err)
			}
			fmt.Println(doc.Data())
		}

	query := client.Collection("users").Where("name", "==", "Frank").Documents(ctx)


	*/

}
