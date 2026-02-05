package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"os"
)

type Reason struct {
	Message string `json:"message"`
}

func main() {

	reasons := readFromFile("./reasons.json")
	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		log.Default().Printf("Received request for /healthz from %s", r.RemoteAddr)
		w.Header().Set("Content-Type", "application/json")
		// Avoid JSON encoding overhead for static response
		w.Write([]byte(`{"status":"UP"}`))
	})

	mux.HandleFunc("GET /no", func(w http.ResponseWriter, r *http.Request) {
		log.Default().Printf("Received request for /no from %s", r.RemoteAddr)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(selectRandomEntry(reasons))
	})

	if err := http.ListenAndServe(defineServerPort(), mux); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func defineServerPort() string {
	var port = os.Getenv("SERVER_PORT")
	if port == "" {
		log.Printf("Using default port 8080")
		port = ":8080"
	} else {
		log.Printf("Overriding default port with %s", port)
		port = fmt.Sprintf(":%s", port)
	}
	return port
}

func readFromFile(pathToFile string) []string {
	log.Printf("Reading reasons from file: %s", pathToFile)
	// Read from file logic goes here
	data, err := os.ReadFile(pathToFile)
	if err != nil {
		log.Printf("Error reading file: %v", err)
		return []string{}
	}
	var strings []string
	if err := json.Unmarshal(data, &strings); err != nil {
		log.Println("Error unmarshalling JSON:", err)
		return []string{}
	}
	log.Println("Current reasons to say no", len(strings))

	return strings
}

func selectRandomEntry(reasons []string) Reason {
	// Select random entry logic goes here
	currentLength := len(reasons)
	if currentLength == 0 {
		return Reason{Message: "Just No!"}
	}
	randomIndex := rand.IntN(currentLength)
	return Reason{Message: reasons[randomIndex]}
}
