package controller

import (
	"fmt"
	"log"
	"song-merger/entities"
	"song-merger/model"
	"strings"
)

// GenerateSongs - Validate song entity before generate the song file.
func GenerateSongs(songs []entities.Song) (string, error) {
	for index, song := range songs {
		song.Artist = strings.TrimSpace(song.Artist)
		song.Name = strings.TrimSpace(song.Name)

		if song.Artist == "" {
			log.Println("[GenerateSong] Error song.Artist == ''")
			message := fmt.Sprintf("%dº song - Artist name can not be empty.", index+1)
			return "", entities.NewBadRequestError(message)
		}

		if song.Name == "" {
			log.Println("[GenerateSong] Error song.Name == ''")
			message := fmt.Sprintf("%dº song - Song name can not be empty.", index+1)
			return "", entities.NewBadRequestError(message)
		}
	}

	return model.GenerateSongs(songs)
}
