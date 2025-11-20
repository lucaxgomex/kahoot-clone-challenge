package main

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/handlers"
)

type Greeting struct {
	Message string `json:"message"`
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	greeting := Greeting{Message: "Hello from Go!"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(greeting)
}

func main() {
	http.HandleFunc("/api/greet", greetHandler)

	// Enable CORS
	http.ListenAndServe(":8081", handlers.CORS()(http.DefaultServeMux))
}
