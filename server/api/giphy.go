package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/michellesauder/go-react-giphy-app/models"

	"github.com/gorilla/mux"
)

func GetGiphs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	searchTerm := vars["search"]

	api_key := "LQQbgB0jgYYH3OrpASCDoewvMJUJXTtI"
	url := "https://api.giphy.com/v1/gifs/search?api_key=" + api_key + "&q=" + searchTerm + "&limit=25&offset=0&rating=g&lang=en&bundle=messaging_non_clips"

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
