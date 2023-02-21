package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	port            = 3060
	defaultUsername = "username"
	defaultPassword = "password"
)

type Response struct {
	Message string `json:"message"`
}

func envOrFallback(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func errorResponse(w *http.ResponseWriter, message string, statusCode int) {
	b, err := json.Marshal(Response{Message: message})
	if err != nil {
		fmt.Println("error:", err)
	}
	http.Error(*w, string(b), statusCode)
}

func main() {
	username := envOrFallback("USERNAME", defaultUsername)
	password := envOrFallback("PASSWORD", defaultPassword)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		workspaceId := r.Header.Get("Plain-Workspace-Id")
		eventType := r.Header.Get("Plain-Event-Type")

		if workspaceId == "" {
			fmt.Println("No Plain-Workspace-Id header found")
			errorResponse(&w, "Bad Request", http.StatusBadRequest)
			return
		}

		if eventType == "" {
			fmt.Println("No Plain-Event-Type header found")
			errorResponse(&w, "Bad Request", http.StatusBadRequest)
			return
		}

		u, p, ok := r.BasicAuth()
		if !ok {
			fmt.Println("No auth credentials were provided in the request")
			errorResponse(&w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if u != username || p != password {
			fmt.Println("Either the username or password do not match")
			errorResponse(&w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			errorResponse(&w, "Unable to read request body", http.StatusBadRequest)
			return
		}

		fmt.Printf("Received %s event from workspace %s: %s\n", eventType, workspaceId, body)

		w.Header().Set("Content-Type", "application/json")
		response := Response{Message: "ok"}
		json.NewEncoder(w).Encode(&response)
	})
	fmt.Printf("Webhook handler running on http://localhost:%d\n", port)
	if err := http.ListenAndServe(fmt.Sprintf("localhost:%d", port), nil); err != http.ErrServerClosed {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println("Server closed")
}
