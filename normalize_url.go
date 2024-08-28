package main

import (
	"net/url"
)

func normalizeURL(rawURL string) (string, error) {
	parsedUrl, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	return parsedUrl.Hostname() + parsedUrl.EscapedPath(), nil
}
