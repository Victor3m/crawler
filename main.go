package main

import (
	"fmt"
)

func main() {
	config, err := configure(getArgs())
	if err != nil {
		fmt.Printf("Error - Configure: %v", err)
	}

	config.crawlPage(config.baseURL.String())

	config.wg.Wait()

	printReport(config.pages, config.baseURL.String())
}
