package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ashutoshbr/GoCRUD/database"
	"github.com/ashutoshbr/GoCRUD/models"
)

func home(w http.ResponseWriter, r *http.Request) {

	myRes := &models.Person{
		Name: "Asta Raven",
		Age:  10,
	}
	data, _ := json.Marshal(myRes)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	log.Println("Listening on port 8000")

	// db connection
	database.Main()

	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}
