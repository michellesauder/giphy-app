package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/michellesauder/go-react-giphy-app/router"
)

func main() {
	r := router.NewRouter()

	// Enable CORS with default options
	cors := handlers.CORS()(r)

	// Start the HTTP server on port 8080
	http.Handle("/", r)
	fmt.Printf("starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", cors))

}
