package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/ashutoshbr/GoCRUD/database"
	"github.com/ashutoshbr/GoCRUD/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Greetings! Welcome to the Homepage. 😁")
}

func read(w http.ResponseWriter, r *http.Request) {
	// db connection
	client := database.Connect()
	coll := client.Database(os.Getenv("DBNAME")).Collection(os.Getenv("COLLNAME"))
	defer client.Disconnect(context.TODO())

	// Safeguard
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		json.NewEncoder(w).Encode("Only GET method allowed!")
		return
	}

	// specify GET JSON format
	filter := bson.D{}
	projection := bson.D{{"uid", 1}, {"name", 1}, {"age", 1}}
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

	// Safeguard
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		json.NewEncoder(w).Encode("Only POST method allowed!")
		return
	}
	p, _ := io.ReadAll(r.Body)
	var temp models.Person
	json.Unmarshal(p, &temp)
	newPerson := models.Person(temp)
	_, err := coll.InsertOne(context.TODO(), newPerson)
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

	// Safeguard
	if r.Method != "PUT" {
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		json.NewEncoder(w).Encode("Only PUT method allowed!")
		return
	}

	p, _ := io.ReadAll(r.Body)
	var temp models.Person
	json.Unmarshal(p, &temp)

	filter := bson.D{{"uid", temp.Uid}}
	update := bson.D{{"$set", bson.D{{"name", temp.Name}}}}
	result, err := coll.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		panic(err)
	}
	fmt.Print("Updated count:", result.ModifiedCount)
	returnStatement := "Updated " + strconv.Itoa(temp.Uid) + " with name " + string(temp.Name)
	json.NewEncoder(w).Encode(returnStatement)
}

func delete(w http.ResponseWriter, r *http.Request) {
	// db connection
	client := database.Connect()
	coll := client.Database(os.Getenv("DBNAME")).Collection(os.Getenv("COLLNAME"))
	defer client.Disconnect(context.TODO())

	// Safeguard
	if r.Method != "DELETE" {
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		json.NewEncoder(w).Encode("Only DELETE method allowed!")
		return
	}

	p, _ := io.ReadAll(r.Body)
	var temp models.Person
	json.Unmarshal(p, &temp)

	filter := bson.D{{"uid", temp.Uid}}
	result, err := coll.DeleteOne(context.TODO(), filter)

	if err != nil {
		panic(err)
	}
	fmt.Print("Deleted count:", result.DeletedCount)
	returnStatement := "Delete item with uid: " + strconv.Itoa(temp.Uid)
	json.NewEncoder(w).Encode(returnStatement)
}
