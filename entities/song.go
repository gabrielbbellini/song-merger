package entities

type Song struct {
	Name        string `json:"name"`
	Artist      string `json:"artist"`
	MusicalTone uint64 `json:"tone"`
}
