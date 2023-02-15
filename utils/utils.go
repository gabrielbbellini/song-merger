package utils

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"song-merger/entities"
)

// createTagRegexp Create a regexp that is looking for a tag by the giving tag name.
func createTagRegexp(tagName string) (*regexp.Regexp, error) {
	tagRegexp, err := regexp.Compile(fmt.Sprintf("<%s>(.|\\n)*?<\\/%s>", tagName, tagName))
	if err != nil {
		log.Println("[createTagRegex] Error Compile")
		return nil, err
	}

	return tagRegexp, nil
}

// ExtractTagFromHTML Extracts the first matched tag and their content from an HTML string.
// tagName is the html tag without the brackets. Example: to extract a "pre" tag and their content, the params will be
// ("pre", "<html><pre>...</pre></html>") witch returns "<pre>...</pre>".
// In case of not found a match tag the return will be an empty string.
func ExtractTagFromHTML(tagName, HTML string) (string, error) {
	tagRegexp, err := createTagRegexp(tagName)
	if err != nil {
		log.Println("[ExtractTagFromHTML] Error createTagRegexp")
		return "", err
	}

	return tagRegexp.FindString(HTML), nil
}

// HandleError Throws the correct http status code and message based on the error type from entities package.
func HandleError(w http.ResponseWriter, err error, tag string) {
	var statusCode int

	switch err.(type) {
	case entities.NotFoundError:
		statusCode = 404

	case entities.BadRequestError:
		statusCode = 400

	default:
		statusCode = 500
	}

	log.Printf("[%s] Error: %s", tag, err.Error())

	w.WriteHeader(statusCode)
	_, _ = w.Write([]byte(err.Error()))
}

// NewErrorFromStatusCode Return the correct error entity based on status code.
func NewErrorFromStatusCode(statusCode int, message string) error {
	switch statusCode {
	case 404:
		return entities.NewNotFoundError(message)

	case 400:
		return entities.NewBadRequestError(message)

	default:
		return entities.NewInternalServerError(message)
	}
}
