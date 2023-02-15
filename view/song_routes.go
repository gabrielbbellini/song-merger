package view

import (
	"encoding/json"
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

	var song entities.Song
	err = json.Unmarshal(b, &song)
	if err != nil {
		utils.HandleError(w, err, "[generateSongs] Error Unmarshal")
		return
	}

	_, err = controller.GenerateSong(song)
	if err != nil {
		utils.HandleError(w, err, "GenerateSong")
		return
	}

	_, _ = w.Write([]byte("Sua música foi gerada com sucesso (Ver pasta \"public\" no projeto)."))
	return
}
