package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type Tests struct {
	name          string
	server        *http.Server
	response      *Weather
	expectedError error
}

func TestGetWeather(t *testing.T) {
	// Creating a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"city": "Denver, Colorado", "weather": "sunny"}`))
	}))
	defer server.Close()

	// Use the server's URL in your test or perform other operations
	// For example, you can make a GET request to the server's URL
	resp, err := http.Get(server.URL)
	if err != nil {
		// Handle the error accordingly
		t.Errorf("Error %v;", err)
		return
	}
	defer resp.Body.Close()

	// Performing tests on the response
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %v; got %v", http.StatusOK, resp.StatusCode)
	}

}
