package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
)

func main() {

	req, err := http.NewRequest(http.MethodGet, "https://www.cifraclub.com.br/harpa-crista/porque-ele-vive/imprimir.html", nil)
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

	r2, err := regexp.Compile("<pre>(.|\\n)*?<\\/pre>")
	if err != nil {
		fmt.Println("Compile", err)
		return
	}

	matched := r2.FindAllString(htmlDomElement, -1)

	b, err = json.Marshal(matched[0])
	if err != nil {
		return
	}

	file, err := os.Create("song.html")
	if err != nil {
		fmt.Println("Create")
		return
	}

	_, err = file.WriteString(matched[0])
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
