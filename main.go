package main

import (
	"log"
	"net/http"
	"song-merger/view"
)

func main() {
	view.SetupRouter()
	fs := http.FileServer(http.Dir("./public/"))
	http.Handle("/", fs)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
