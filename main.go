package main

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net/http"
	"os"
)

func get_random_joke() string {
	// Read file
	data, err := os.ReadFile("jokes.json")
	if err != nil {
		panic(err)
	}

	// Decode jokes into a slice of strings
	var jokes []string
	err = json.Unmarshal(data, &jokes)
	if err != nil {
		panic(err)
	}

	return jokes[rand.IntN(len(jokes))]
}

func main() {
	// Serve static files out of the ./static folder
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.HandleFunc("/api/joke", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// Create a proper response struct and marshal it to JSON
		response := map[string]string{
			"message": get_random_joke(),
		}

		json.NewEncoder(w).Encode(response)
	})

	// Use railway's port env if available
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server running on localhost:%s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Error while running server: ", err)
	}
}
