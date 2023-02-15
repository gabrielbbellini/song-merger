package model

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"song-merger/entities"
	"song-merger/store"
	"song-merger/utils"
	"strconv"
)

func requestSong(songURL string) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodGet, songURL, nil)
	if err != nil {
		fmt.Println("[requestSong] Error NewRequest")
		return nil, err
	}

	client := &http.Client{}
	return client.Do(request)
}

func getSongHTMLPage(song entities.Song) (string, error) {
	url := fmt.Sprintf("https://www.cifraclub.com.br/%s/%s/imprimir.html#key=%s",
		song.Artist,
		song.Name,
		strconv.FormatUint(song.MusicalTone, 10),
	)

	response, err := requestSong(url)
	if err != nil {
		log.Println("[GenerateSong] Error requestSong")
		return "", err
	}
	if response.StatusCode != http.StatusOK {
		log.Println("[GenerateSong] Error response.StatusCode != http.StatusOK")
		return "", utils.NewErrorFromStatusCode(response.StatusCode, "Erro inesperado ao consultar cifra")
	}

	b, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("[getSong] Error ReadAll")
		return "", err
	}

	return string(b), nil
}

//GenerateSong - Create the merged songs file by doing requests and get the song score.
func GenerateSong(song entities.Song) (string, error) {
	html, err := getSongHTMLPage(song)
	if err != nil {
		log.Println("[GenerateSong] Error getSongHTMLPage")
		return "", err
	}

	score, err := utils.ExtractTagFromHTML("pre", html)
	if err != nil {
		log.Println("[getSong] Error ExtractTagFromHTML")
		return "", err
	}
	if score == "" {
		log.Println("[getSong] Error ExtractTagFromHTML")
		return "", entities.NewNotFoundError(fmt.Sprintf("Cifra da música \"%s\" não foi encontrada.", song.Name))
	}

	err = createSongFile(fmt.Sprintf(
		"%s-%s-%d.html",
		song.Name,
		song.Artist,
		song.MusicalTone,
	), score)
	if err != nil {
		log.Println("[getSong] Error createSongPage")
		return "", err
	}

	return "", nil
}

func generateSongHTMLTemplate(score string) string {
	return fmt.Sprintf(`
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
}

func generateSongFile(fileName, html string) error {
	err := store.Manager().CreateFile(fileName)
	if err != nil {
		log.Println("[createSongPage] Error CreateFile")
		return err
	}

	err = store.Manager().WriteStringFile(fileName, html)
	if err != nil {
		log.Println("[createSongPage] Error WriteStringFile")
		return err
	}

	return nil
}

func createSongFile(fileName, score string) error {
	html := generateSongHTMLTemplate(score)
	return generateSongFile(fileName, html)
}
