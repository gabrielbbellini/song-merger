package entities

type Song struct {
	Name   string `json:"name"`
	Artist string `json:"artist"`
	Tone   uint64 `json:"tone"`
	Tabs   bool   `json:"tabs"`
}
