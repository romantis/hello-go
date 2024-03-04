package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Message struct {
	Greeting string `json:"greeting"`
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	message := Message{"Hello, Odyn!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

func main() {
	http.HandleFunc("/", helloWorld)
	log.Println("Server starting on port 8081...")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
