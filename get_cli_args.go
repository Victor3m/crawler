package main

import (
	"fmt"
	"os"
	"strconv"
)

func getArgs() (string, int, int) {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(args) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	fmt.Printf("starting crawl of: %s\n", args[0])
	URL := args[0]
	maxConcurrency, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Printf("Error - Argument 2 is not a valid int: %v\n")
		os.Exit(1)
	}
	maxPages, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Printf("Error - Argument 3 is not a valid int: %v\n")
		os.Exit(1)
	}
	return URL, maxConcurrency, maxPages
}
