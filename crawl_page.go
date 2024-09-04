package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	fmt.Println(rawCurrentURL)

	baseURL, _ := url.Parse(rawBaseURL)
	currentURL, _ := url.Parse(rawCurrentURL)

	if currentURL.Hostname() != baseURL.Hostname() {
		return
	}

	normalURL, _ := normalizeURL(rawCurrentURL)

	_, ok := pages[normalURL]

	if ok {
		pages[normalURL]++
		return
	} else {
		pages[normalURL] = 1
	}

	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("There was an error getting the HTML: %v", err)
	}

	//fmt.Println(html)

	urls, err := getURLsFromHTML(html, rawBaseURL)
	if err != nil {
		fmt.Printf("There was an error getting the URLs from the Html: %v", err)
	}

	for _, url := range urls {
		crawlPage(rawBaseURL, url, pages)
	}

	return
}
