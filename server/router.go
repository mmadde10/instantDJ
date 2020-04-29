package main

import (
	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/autenticate", AuthenticateUser).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/callback", CompleteAuth).Methods("GET", "OPTIONS")

	//Track routes
	router.HandleFunc("/api/tracks/{id}", GetTrack).Methods("GET", "OPTIONS")

	//Search Route
	router.HandleFunc("/api/search/{query}", GetSearchResults).Methods("GET", "OPTIONS")

	return router
}
