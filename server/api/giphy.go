package api

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/michellesauder/go-react-giphy-app/models"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func GetGiphs(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	ctx = context.WithValue(ctx, "env", os.Environ())

	// Access environment variables
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		http.Error(w, "API_KEY not found", http.StatusInternalServerError)
		return
	}

	ctx = context.WithValue(ctx, "apiKey", apiKey)

	vars := mux.Vars(r)
	searchTerm := vars["search"]
	ctx = context.WithValue(ctx, "searchTerm", searchTerm)

	url := "https://api.giphy.com/v1/gifs/search?api_key=" + apiKey + "&q=" + searchTerm + "&limit=25&offset=0&rating=g&lang=en&bundle=messaging_non_clips"

	resp, getErr := http.Get(url)
	if getErr != nil {
		http.Error(w, getErr.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		http.Error(w, readErr.Error(), http.StatusInternalServerError)
		return
	}

	// Unmarshal the JSON response into an instance of the ResponseData struct.
	var responseData models.ResponseData
	if err := json.Unmarshal(body, &responseData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Serialize the responseData to JSON and send it as the HTTP response.
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(responseData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
