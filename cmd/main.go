package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getHome)
	mux.HandleFunc("/create", create)
	mux.HandleFunc("/read", read)
	mux.HandleFunc("/update", update)
	mux.HandleFunc("/delete", delete)
	log.Println("Listening on port 8000")

	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}
