package view

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"song-merger/controller"
	"song-merger/entities"
)

func SetSongRoutes() {
	http.HandleFunc("/songs", generateSongs)
}

func generateSongs(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("[generateSongs] Error ReadAll")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var songRequest entities.SongRequest
	err = json.Unmarshal(b, &songRequest)
	if err != nil {
		fmt.Println("[generateSongs] Error ReadAll")
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	_, exception := controller.GenerateSong(songRequest)
	if exception != nil {
		fmt.Println("[generateSongs] Error ReadAll")
		http.Error(w, exception.Message, exception.Code)
		return
	}

	http.Redirect(w, r, "http://localhost:8000/song.html", http.StatusSeeOther)
	return
}
