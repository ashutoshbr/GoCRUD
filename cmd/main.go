package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	// mux.HandleFunc("/", GetHome)
	mux.HandleFunc("/", GetHome)
	log.Println("Listening on port 8000")

	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}
