package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", GetHome)
	mux.HandleFunc("/create", Create)
	mux.HandleFunc("/update", Update)
	log.Println("Listening on port 8000")

	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}
