package view

import "github.com/gorilla/mux"

// SetupRouter - Create all routes handlers
func SetupRouter(router *mux.Router) {
	SetSongRoutes(router)
}
