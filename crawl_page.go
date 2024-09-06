package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {

	cfg.concurrencyControl <- struct{}{}

	defer func() {
		cfg.wg.Done()
		<-cfg.concurrencyControl
	}()

	cfg.mu.Lock()
	numPages := len(cfg.pages)
	cfg.mu.Unlock()

	if numPages >= cfg.maxPages {
		return
	}

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
	}

	if currentURL.Hostname() != cfg.baseURL.Hostname() {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - normalizedURL: %v", err)
	}

	isFirst := cfg.addPageVisit(normalizedURL)

	if isFirst {
	} else {
		return
	}

	fmt.Printf("Crawling %s\n", rawCurrentURL)

	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - getHTML: %v", err)
	}

	urls, err := getURLsFromHTML(html, cfg.baseURL.String())
	if err != nil {
		fmt.Printf("Error - getURLsFromHTML: %v", err)
	}

	for _, nextURL := range urls {
		cfg.wg.Add(1)
		go cfg.crawlPage(nextURL)
	}

	return
}
