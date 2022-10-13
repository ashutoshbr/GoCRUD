package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Namaste"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	log.Println("Listening on port 8000")
	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}
