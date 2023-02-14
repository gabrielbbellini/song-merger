package controller

import (
	"log"
	"song-merger/entities"
	utils "song-merger/exception"
	"song-merger/model"
	"strings"
)

func GenerateSong(song entities.SongRequest) (string, *utils.Exception) {
	song.ArtistName = strings.TrimSpace(song.ArtistName)
	song.Name = strings.TrimSpace(song.Name)

	if song.ArtistName == "" {
		log.Println("[GenerateSong] Error song.ArtistName == \"\"")
		return "", utils.NewException("Artist name can not be empty.", 400)
	}

	if song.Name == "" {
		log.Println("[GenerateSong] Error song.Name == \"\"")
		return "", utils.NewException("Song name can not be empty.", 400)
	}

	return model.GenerateSong(song)
}
