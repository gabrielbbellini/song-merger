package model

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"song-merger/entities"
	utils "song-merger/exception"
	"strconv"
)

func requestSong(songURL string) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodGet, songURL, nil)
	if err != nil {
		fmt.Println("NewRequest", err)
		return nil, err
	}

	client := &http.Client{}
	return client.Do(request)
}

func newSongURL(song entities.SongRequest) string {
	return fmt.Sprintf("https://www.cifraclub.com.br/%s/%s/imprimir.html?key=%s",
		song.ArtistName,
		song.Name,
		strconv.FormatUint(song.MusicalTone, 10),
	)
}

func GenerateSong(song entities.SongRequest) (string, *utils.Exception) {
	url := newSongURL(song)
	response, err := requestSong(url)
	if err != nil {
		log.Println("[GenerateSong] Error requestSong")
		return "", utils.NewException(err.Error(), 500)
	}
	if response.StatusCode != http.StatusOK {
		fmt.Println("[GenerateSong] Error response.StatusCode != http.StatusOK")
		return "", utils.NewException(
			fmt.Sprintf("A requisição para %s falhou. Error: %s", url, response.Status),
			500,
		)
	}

	b, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("[getSong] Error ReadAll", err)
		return "", utils.NewException(err.Error(), 500)
	}

	htmlElement := string(b)
	score, exception := extractPreTagFromHTML(htmlElement)
	if err != nil {
		fmt.Println("[getSong] Error extractPreTagFromHTML", err)
		return "", exception
	}

	exception = createSongPage("song.html", score)
	if exception != nil {
		fmt.Println("[getSong] Error createSongPage", err)
		return "", exception
	}

	return "", nil
}

func createSongPage(filename, score string) *utils.Exception {
	// Put score into HTML output
	//language=HTML
	htmlPage := fmt.Sprintf(`
		<html>
			<head>
				<title>
					SONG GENERATOR
				</title>
			</head>
			<main>
				%s
			</main>
		</html>
	`, score)

	// Create and write html to output file
	file, err := os.Create("./public/" + filename)
	if err != nil {
		log.Println("[createSongPage] Error Create")
		return nil
	}

	_, err = file.WriteString(htmlPage)
	if err != nil {
		log.Println("[createSongPage] Error WriteString")
		return utils.NewException(err.Error(), 500)
	}

	err = file.Close()
	if err != nil {
		fmt.Println("[createSongPage] Error Close")
		return utils.NewException(err.Error(), 500)
	}

	return nil
}

// extractPreTagFromHTML Extract only the score contents from song's HTML
func extractPreTagFromHTML(htmlElement string) (string, *utils.Exception) {
	r2, err := regexp.Compile("<pre>(.|\\n)*?<\\/pre>")
	if err != nil {
		fmt.Println("[extractPreTagFromHTML] Error Compile", err)
		return "", utils.NewException(err.Error(), 500)
	}
	matched := r2.FindAllString(htmlElement, -1)

	if len(matched) < 1 {
		fmt.Println("[extractPreTagFromHTML] Error len(matched) < 1")
		return "", utils.NewException("Score not found", 500)
	}

	return matched[0], nil
}
