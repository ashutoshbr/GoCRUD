package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/ashutoshbr/GoCRUD/database"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	// db connection
	client := database.Connect()
	coll := client.Database(os.Getenv("DBNAME")).Collection(os.Getenv("COLLNAME"))
	defer client.Disconnect(context.TODO())

	// specify GET JSON format
	filter := bson.D{}
	projection := bson.D{{"name", 1}, {"age", 1}, {"_id", 0}}
	opts := options.Find().SetProjection(projection)

	cursor, err := coll.Find(context.TODO(), filter, opts)
	if err != nil {
		panic(err)
	}

	var results []bson.D
	err = cursor.All(context.TODO(), &results)
	for _, result := range results {
		fmt.Println(result)
	}
}

// func Create(w http.ResponseWriter, r *http.Request) {
// 	// db connection
// 	// client, ctx, cancel := database.Main()
// 	var client, ctx, cancel, err = database.Connect(URI)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer database.Close(client, ctx, cancel)
// 	collection := client.Database(os.Getenv("DBNAME")).Collection(os.Getenv("COLLNAME"))

// 	person1 := models.Person{"Foo Bar", 20}
// 	_, err = collection.InsertOne(ctx, person1)
// 	if err != nil {
// 		panic(err)
// 	}
// }
