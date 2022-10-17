package main

import (
	"context"
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

func Create(w http.ResponseWriter, r *http.Request) {
	// db connection
	client := database.Connect()
	coll := client.Database(os.Getenv("DBNAME")).Collection(os.Getenv("COLLNAME"))
	defer client.Disconnect(context.TODO())

	person1 := models.Person{"Abc Xyz", 5}
	_, err := coll.InsertOne(context.TODO(), person1)
	if err != nil {
		panic(err)
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	// db connection
	client := database.Connect()
	coll := client.Database(os.Getenv("DBNAME")).Collection(os.Getenv("COLLNAME"))
	defer client.Disconnect(context.TODO())

	filter := bson.D{{"age", 5}}
	update := bson.D{{"$set", bson.D{{"age", 50}}}}
	result, err := coll.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		panic(err)
	}
	fmt.Print("Updated count:", result.ModifiedCount)
}
