package main

import (
	"log"
	"net/http"
	"song-merger/view"
	"time"

	"github.com/gorilla/mux"
)

const SERVER_URL string = "127.0.0.1:8000"

// Apply the rules to the api router
func apiMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		handler.ServeHTTP(w, r)
	})
}

func main() {
	router := mux.NewRouter()

	// Setup api router.
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.Use(apiMiddleware)
	view.SetupRouter(apiRouter)

	server := &http.Server{
		Handler:      router,
		Addr:         SERVER_URL,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Listening on %s...", SERVER_URL)

	log.Fatal(server.ListenAndServe())
}
