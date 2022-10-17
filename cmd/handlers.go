package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/ashutoshbr/GoCRUD/database"
	"github.com/ashutoshbr/GoCRUD/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Greetings! Welcome to the Homepage.")
}

func read(w http.ResponseWriter, r *http.Request) {
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
	cursor.All(context.TODO(), &results)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func create(w http.ResponseWriter, r *http.Request) {
	// db connection
	client := database.Connect()
	coll := client.Database(os.Getenv("DBNAME")).Collection(os.Getenv("COLLNAME"))
	defer client.Disconnect(context.TODO())

	if r.Method != "POST" {
		json.NewEncoder(w).Encode("Only POST method allowed!")
		panic("Only POST method allowed!")
	}
	p, _ := io.ReadAll(r.Body)
	var temp models.Person
	json.Unmarshal(p, &temp)
	person1 := models.Person(temp)
	_, err := coll.InsertOne(context.TODO(), person1)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(temp)
}

func update(w http.ResponseWriter, r *http.Request) {
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

func delete(w http.ResponseWriter, r *http.Request) {
	// db connection
	client := database.Connect()
	coll := client.Database(os.Getenv("DBNAME")).Collection(os.Getenv("COLLNAME"))
	defer client.Disconnect(context.TODO())

	filter := bson.D{{"age", 5}}
	result, err := coll.DeleteOne(context.TODO(), filter)

	if err != nil {
		panic(err)
	}
	fmt.Print("Deleted count:", result.DeletedCount)
}
