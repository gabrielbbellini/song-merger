package model

import (
	"bytes"
	"fmt"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/google/uuid"
	"html/template"
	"io"
	"log"
	"net/http"
	"path"
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

// GenerateSongs - Create the merged songs file by doing requests and get the song score.
func GenerateSongs(songs []entities.Song) (string, error) {
	var songsTemplateData []entities.SongTemplateData
	for _, song := range songs {
		songScore, err := getSongScore(song)
		if err != nil {
			log.Println("[GenerateSongs] Error generateSong")
			return "", err
		}

		songTemplateData := entities.SongTemplateData{
			Song:        song,
			HTMLContent: template.HTML(songScore),
		}

		songsTemplateData = append(songsTemplateData, songTemplateData)
	}

	return generateSongsHTMLTemplate(songsTemplateData)
}

func getSongScore(song entities.Song) (string, error) {
	html, err := getSongHTMLPage(song)
	if err != nil {
		log.Println("[generateSong] Error getSongHTMLPage")
		return "", err
	}

	score, err := utils.ExtractTagFromHTML("pre", html)
	if err != nil {
		log.Println("[generateSong] Error ExtractTagFromHTML")
		return "", err
	}
	if score == "" {
		log.Println("[generateSong] Error ExtractTagFromHTML")
		return "", entities.NewNotFoundError(fmt.Sprintf("Cifra da música \"%s\" não foi encontrada.", song.Name))
	}

	return score, nil
}

func generateSongsHTMLTemplate(songsTemplateData []entities.SongTemplateData) (string, error) {
	var templ *template.Template
	var err error

	if templ, err = template.ParseFiles("./html_song_model/html_song_model.html"); err != nil {
		log.Println("[generateSongsHTMLTemplate] Error ParseFiles")
		return "", err
	}
	pdfGenerator, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Println("[generateSongsHTMLTemplate] Error NewPDFGenerator")
		return "", err
	}

	var songTemplate bytes.Buffer
	if err = templ.Execute(&songTemplate, songsTemplateData); err != nil {
		log.Println("[generateSongsHTMLTemplate] Error Execute")
		return "", err
	}

	page := wkhtmltopdf.NewPageReader(bytes.NewReader(songTemplate.Bytes()))

	page.EnableLocalFileAccess.Set(true)

	pdfGenerator.AddPage(page)

	pdfGenerator.MarginLeft.Set(0)
	pdfGenerator.MarginRight.Set(0)
	pdfGenerator.Dpi.Set(300)
	pdfGenerator.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfGenerator.Orientation.Set(wkhtmltopdf.OrientationLandscape)

	err = pdfGenerator.Create()
	if err != nil {
		log.Println("[generateSongsHTMLTemplate] Error Create")
		return "", err
	}

	storageManager := store.Manager()

	rootPath := storageManager.RootPath
	filename := uuid.New().String() + ".pdf"

	err = pdfGenerator.WriteFile(path.Join(rootPath, filename))
	if err != nil {
		log.Println("[generateSongsHTMLTemplate] Error WriteFile")
		return "", err
	}

	return filename, nil
}
