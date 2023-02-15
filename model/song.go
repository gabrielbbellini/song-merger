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
	"strings"
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
		strconv.FormatUint(song.Tone, 10),
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

	err = createSongFile(song, score)
	if err != nil {
		log.Println("[getSong] Error createSongPage")
		return "", err
	}

	return "", nil
}

func generateSongHTMLTemplate(song entities.Song, score string) string {
	return fmt.Sprintf(`
		<html>
			<head>
				<title>
					SONG GENERATOR
				</title>
			</head>
			<main>
				<h1>%s</h1>
				<h2>%s</h2>
				%s
			</main>
		</html>
	`, strings.Replace(song.Name, "-", " ", -1),
		strings.Replace(song.Artist, "-", " ", -1),
		score)
}

func generateSongFile(song entities.Song, html string) error {
	fileName := fmt.Sprintf(
		"%s-%s-%d.html",
		song.Name,
		song.Artist,
		song.Tone,
	)

	storeManager := store.Manager()
	err := storeManager.CreateFile(fileName)
	if err != nil {
		log.Println("[createSongPage] Error CreateFile")
		return err
	}

	err = storeManager.WriteStringFile(fileName, html)
	if err != nil {
		log.Println("[createSongPage] Error WriteStringFile")
		return err
	}

	return nil
}

func createSongFile(song entities.Song, score string) error {
	html := generateSongHTMLTemplate(song, score)
	return generateSongFile(song, html)
}
