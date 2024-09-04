package main

import (
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputBody string
		inputUrl  string
		expected  []string
	}{
		{
			name: "absolute and relative URLs",
			inputBody: `
			<html>
				<body>
					<a href="/path/one">
						<span>Boot.dev</span>
					</a>
					<a href="https://other.com/path/one">
						<span>Boot.dev</span>
					</a>
				</body>
			</html>
			`,
			inputUrl: "https://blog.boot.dev",
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name: "lots of links",
			inputBody: `
				<!DOCTYPE html>
				<html lang="en">
				<head>
						<meta charset="UTF-8">
						<meta name="viewport" content="width=device-width, initial-scale=1.0">
						<title>Anchor Tags and Links</title>
				</head>
				<body>
						<h1>Welcome to the Anchor Tags and Links Page</h1>
						
						<h2>Absolute Links</h2>
						<ul>
								<li><a href="https://www.example.com">Example Domain</a></li>
								<li><a href="https://www.wikipedia.org">Wikipedia</a></li>
								<li><a href="https://www.github.com">GitHub</a></li>
								<li><a href="https://www.stackoverflow.com">Stack Overflow</a></li>
								<li><a href="https://www.google.com">Google</a></li>
						</ul>

						<h2>Relative Links</h2>
						<ul>
								<li><a href="/about.html">About Us</a></li>
								<li><a href="/contact.html">Contact</a></li>
								<li><a href="../services.html">Our Services</a></li>
								<li><a href="../../index.html">Home</a></li>
								<li><a href="/folder/page.html">Nested Page</a></li>
						</ul>

						<h2>Mixed Links</h2>
						<ul>
								<li><a href="https://www.example.com/folder/page.html">Absolute with Path</a></li>
								<li><a href="https://www.example.com/long/path/to/resource">Long Absolute Path</a></li>
								<li><a href="/short.html">Short Relative Path</a></li>
						</ul>
				</body>
				</html>
			`,
			inputUrl: "https://zzzcode.ai/resource/help",
			expected: []string{"https://www.example.com", "https://www.wikipedia.org", "https://www.github.com", "https://www.stackoverflow.com", "https://www.google.com", "https://zzzcode.ai/resource/help/about.html", "https://zzzcode.ai/resource/help/contact.html", "https://zzzcode.ai/resource/services.html", "https://zzzcode.ai/index.html", "https://zzzcode.ai/resource/help/folder/page.html", "https://www.example.com/folder/page.html", "https://www.example.com/long/path/to/resource", "https://zzzcode.ai/resource/help/short.html"},
		},
		{
			name: "no links",
			inputBody: `
				<html lang="en">
					<body>
						<h1>NO Anchor Tags here</h1>
					</body>
				</html>
			`,
			inputUrl: "https://google.com/",
			expected: nil,
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputUrl)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - '%s' FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
