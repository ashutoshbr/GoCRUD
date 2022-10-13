package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type SimpleResponse struct {
	Msg string `json:"msg"`
}

func home(w http.ResponseWriter, r *http.Request) {

	myRes := &SimpleResponse{
		Msg: "Namaste",
	}
	data, _ := json.Marshal(myRes)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	log.Println("Listening on port 8000")
	// check if env variables are loaded
	fmt.Println(os.Getenv("MONGOUSER"))
	err = http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}
