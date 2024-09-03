package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	urls, err := getURLsFromHTML(`
			<html>
				<body>
					<a name="link" href="/path/one">
						<span>Boot.dev</span>
					</a>
					<a href="https://other.com/path/one">
						<span>Boot.dev</span>
					</a>
				</body>
			</html>
	`, "boot.dev")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	for _, url := range urls {
		fmt.Println(url)
	}
}
