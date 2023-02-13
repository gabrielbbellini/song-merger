package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
)

func main() {

	// Get score for a given song:
	artistName := "harpa-crista"
	songName := "porque-ele-vive"

	// this link can be generated on the fly; just format the artist and song name accordingly
	songScoreLink := fmt.Sprintf("https://www.cifraclub.com.br/%s/%s/imprimir.html",
		artistName,
		songName,
	)

	req, err := http.NewRequest(http.MethodGet, songScoreLink, nil)
	if err != nil {
		fmt.Println("NewRequest", err)
		return
	}

	req.Header.Add("Content-Type", "text/html")
	req.Header.Add("Accept-Charset", "utf-8")

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Do")
		return
	}

	b, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("ReadAll", err)
		return
	}

	htmlDomElement := string(b)

	// Extract only the score contents from song's HTML
	r2, err := regexp.Compile("<pre>(.|\\n)*?<\\/pre>")
	if err != nil {
		fmt.Println("Compile", err)
		return
	}

	matched := r2.FindAllString(htmlDomElement, -1)

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
	`, matched[0])

	// Create and write html to output file
	file, err := os.Create("song.html")
	if err != nil {
		fmt.Println("Create")
		return
	}

	_, err = file.WriteString(htmlPage)
	if err != nil {
		log.Println("WriteFile")
		return
	}

	err = file.Close()
	if err != nil {
		fmt.Println("Write")
		return
	}
}
