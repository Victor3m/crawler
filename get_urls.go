package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	depth := 0
	reader := strings.NewReader(htmlBody)
	z := html.NewTokenizer(reader)
	var urls []string
	inputURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, err
	}

	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return urls, nil
		case html.TextToken:
			if depth > 0 {
				keys, _, _ := z.TagAttr()
				for _, key := range keys {
					fmt.Println(key)
				}
			}
		case html.StartTagToken, html.EndTagToken:
			tn, _ := z.TagName()
			if len(tn) == 1 && tn[0] == 'a' {
				if tt == html.StartTagToken {
					depth++
					for {
						key, value, moreAttr := z.TagAttr()
						if string(key) == "href" {
							parsedUrl, err := url.Parse(string(value))
							if err != nil {
								return nil, err
							}
							if parsedUrl.Host == "" {
								parsedUrl = inputURL.JoinPath(parsedUrl.Path)
							}
							urls = append(urls, parsedUrl.String())
						}
						if !moreAttr {
							break
						}
					}
				} else {
					depth--
				}
			}
		}
	}
}
