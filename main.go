package main

import (
	"encoding/json"
	"log"
	"net/http"
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
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	log.Println("Listening on port 8000")
	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}
