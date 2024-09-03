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
							if parsedUrl.Scheme == "" {
								parsedUrl.Scheme = "https"
							}
							if parsedUrl.Host == "" {
								parsedUrl.Host = rawBaseURL
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
