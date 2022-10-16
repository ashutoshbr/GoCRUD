package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ashutoshbr/GoCRUD/database"
	"github.com/ashutoshbr/GoCRUD/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	// db connection
	client, ctx, cancel := database.Main()
	defer database.Close(client, ctx, cancel)
	collection := client.Database(os.Getenv("DBNAME")).Collection(os.Getenv("COLLNAME"))

	filter := bson.D{}
	projection := bson.D{{"name", 1}, {"age", 1}, {"_id", 0}}
	opts := options.Find().SetProjection(projection)

	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		panic(err)
	}
	var results []bson.D
	err = cursor.All(ctx, &results)
	for _, result := range results {
		fmt.Println(result)
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	// db connection
	client, ctx, cancel := database.Main()
	defer database.Close(client, ctx, cancel)
	collection := client.Database(os.Getenv("DBNAME")).Collection(os.Getenv("COLLNAME"))

	person1 := models.Person{"Foo Bar", 20}
	_, err := collection.InsertOne(ctx, person1)
	if err != nil {
		panic(err)
	}
}
