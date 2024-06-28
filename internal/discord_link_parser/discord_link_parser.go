package discord_link_parser

import (
	"io"
	"net/http"
)

const (
	GIST_LINK = "https://gist.githubusercontent.com/Insadem/6e6c7971d1c7828fb44b182e6fd12ca0/raw"
)

var cachedDiscordLink string

func parseGist() (string, error) {
	response, err := http.Get(GIST_LINK)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	result, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(result), err
}

func DiscordLink() string {
	result, err := parseGist()
	if err != nil {
		return cachedDiscordLink
	}

	cachedDiscordLink = result
	return result
}
