package view

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"song-merger/controller"
	"song-merger/entities"
	"song-merger/utils"
)

// SetSongRoutes - Create all song entity routes.
func SetSongRoutes(router *mux.Router) {
	router.HandleFunc("/songs", generateSongs).Methods(http.MethodPost)
}

// generateSongs - Return the PDF that contains the merged songs.
func generateSongs(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("[generateSongs] Error ReadAll")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var songs []entities.Song
	err = json.Unmarshal(b, &songs)
	if err != nil {
		log.Println("[generateSongs] Error Unmarshal")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	filename, err := controller.GenerateSongs(songs)
	if err != nil {
		utils.HandleError(w, err, "GenerateSong")
		return
	}

	path := "http://127.0.0.1:8000/public/" + filename
	_, _ = w.Write([]byte(fmt.Sprintf("Sua música foi gerada com sucesso: %s", path)))
	return
}
