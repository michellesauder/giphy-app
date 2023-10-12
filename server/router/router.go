package router

import (
	"my-giphy-app/api"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/giphs", api.GetGiphs).Methods("GET")
	return r
}
