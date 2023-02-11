package model

import (
	"encoding/json"
	"io"
	"net/http"
)

type SongRepository struct {
}

func NewSongRepository(db) {

}

func (s SongRepository) GetSong(songName string) string {
	response, err := http.Get("")
	if err != nil {

	}

	b, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}

	var songPage string
	err = json.Unmarshal(b, &songPage)
	if err != nil {
		return
	}

	return songPage
}
