package router

import (
	"github.com/michellesauder/go-react-giphy-app/api"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/giphs/{search}", api.GetGiphs).Methods("GET")
	return r
}
