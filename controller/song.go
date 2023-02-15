package controller

import (
	"log"
	"song-merger/entities"
	"song-merger/model"
	"strings"
)

// GenerateSong - Validate song entity before generate the song file.
func GenerateSong(song entities.Song) (string, error) {
	song.Artist = strings.TrimSpace(song.Artist)
	song.Name = strings.TrimSpace(song.Name)

	if song.Artist == "" {
		log.Println("[GenerateSong] Error song.Artist == ''")
		return "", entities.NewBadRequestError("Artist name can not be empty.")
	}

	if song.Name == "" {
		log.Println("[GenerateSong] Error song.Name == ''")
		return "", entities.NewBadRequestError("Song name can not be empty.")
	}

	return model.GenerateSong(song)
}
