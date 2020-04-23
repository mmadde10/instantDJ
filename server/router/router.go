package router

import (
	"go-instant-dj/instantDJ/server/middleware"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/autenticate", middleware.AuthenticateUser).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/callback", middleware.CompleteAuth).Methods("GET", "OPTIONS")

	router.HandleFunc("/api/tracks/{id}", middleware.GetTrack).Methods("GET", "OPTIONS")

	return router
}
