package main

import (
	"fmt"
	"os"
)

func getArgs() string {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(args) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	fmt.Printf("starting crawl of: %s\n", args[0])
	return args[0]
}
