package entities

type SongRequest struct {
	Name        string `json:"name"`
	ArtistName  string `json:"artistName"`
	MusicalTone uint64 `json:"musicalTone"`
}

type SongResponse struct {
	HTMLElement string `json:"htmlElement"`
}
