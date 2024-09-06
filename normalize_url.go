package main

import (
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	parsedUrl, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(strings.ToLower(parsedUrl.Host+parsedUrl.Path), "/"), nil
}
