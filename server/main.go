package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Image struct {
	Height   string `json:"height"`
	Width    string `json:"width"`
	Size     string `json:"size"`
	URL      string `json:"url"`
	MP4Size  string `json:"mp4_size"`
	MP4      string `json:"mp4"`
	WebPSize string `json:"webp_size"`
	WebP     string `json:"webp"`
	Frames   string `json:"frames"`
	Hash     string `json:"hash"`
}

type DataItem struct {
	Type             string `json:"type"`
	ID               string `json:"id"`
	URL              string `json:"url"`
	Slug             string `json:"slug"`
	BitlyGifURL      string `json:"bitly_gif_url"`
	BitlyURL         string `json:"bitly_url"`
	EmbedURL         string `json:"embed_url"`
	Username         string `json:"username"`
	Source           string `json:"source"`
	Title            string `json:"title"`
	Rating           string `json:"rating"`
	ContentURL       string `json:"content_url"`
	SourceTLD        string `json:"source_tld"`
	SourcePostURL    string `json:"source_post_url"`
	IsSticker        int    `json:"is_sticker"`
	ImportDatetime   string `json:"import_datetime"`
	TrendingDatetime string `json:"trending_datetime"`
	Images           struct {
		Original               Image `json:"original"`
		FixedHeight            Image `json:"fixed_height"`
		FixedHeightDownsampled Image `json:"fixed_height_downsampled"`
		FixedWidth             Image `json:"fixed_width"`
		FixedWidthDownsampled  Image `json:"fixed_width_downsampled"`
	} `json:"images"`
}

type Pagination struct {
	TotalCount int `json:"total_count"`
	Count      int `json:"count"`
	Offset     int `json:"offset"`
}

type Meta struct {
	Status     int    `json:"status"`
	Message    string `json:"msg"`
	ResponseID string `json:"response_id"`
}

type ResponseData struct {
	Data       []DataItem `json:"data"`
	Pagination Pagination `json:"pagination"`
	Meta       Meta       `json:"meta"`
}

//functions

func getGiphs(w http.ResponseWriter, r *http.Request) {
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
	var responseData ResponseData
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

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/giphs/{search}", getGiphs).Methods("GET")

	// Enable CORS with default options
	cors := handlers.CORS()(r)

	// Start the HTTP server on port 8080
	http.Handle("/", r)
	fmt.Printf("starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", cors))

}
