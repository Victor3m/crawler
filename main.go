package main

import "fmt"

func main() {
	pages := map[string]int{}
	crawlPage(getArgs(), getArgs(), pages)

	for key, value := range pages {
		fmt.Printf("URL: %s -- Mentions: %d\n", key, value)
	}
}
