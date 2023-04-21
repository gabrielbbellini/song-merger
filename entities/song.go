package entities

import "html/template"

// Song entity that contains all information about the song to generate
type Song struct {
	Name   string `json:"name"`
	Artist string `json:"artist"`
	Tone   uint64 `json:"tone"`
	Tabs   bool   `json:"tabs"`
}

// SongTemplate entity that is used to generate the html file.
type SongTemplate struct {
	Song        Song
	HTMLContent template.HTML
}
