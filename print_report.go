package main

import (
	"fmt"
	"sort"
	"strings"
)

type Pages struct {
	URL   string
	Count int
}

func sortPages(p []*Pages) {
	sort.Slice(p, func(i, j int) bool {
		return p[i].Count > p[j].Count
	})
}

func printReport(pages map[string]int, baseURL string) {
	line := strings.Repeat("=", len(baseURL)+13)
	fmt.Println(line)
	fmt.Printf(" REPORT for %s\n", baseURL)
	fmt.Println(line)

	pagesSlice := []*Pages{}

	for url, count := range pages {
		tempPage := &Pages{
			URL:   url,
			Count: count,
		}

		pagesSlice = append(pagesSlice, tempPage)
	}

	sortPages(pagesSlice)

	for _, page := range pagesSlice {
		fmt.Printf("Found %d internal links to %s\n", page.Count, page.URL)
	}
}
