package main

import "testing"

func TestCrawlPage(t *testing.T) {
	tests := []struct {
		name       string
		rawBaseURL string
		pages      map[string]int
		expected   string
	}{
		{
			name:       "wagslane blog site",
			rawBaseURL: "https://wagslane.dev/",
			pages:      map[string]int{},
			expected:   "",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			crawlPage(tc.rawBaseURL, tc.rawBaseURL, tc.pages)
			if len(tc.pages) == 0 {
				t.Errorf("Test %v - '%s' FAIL: empty map!", i, tc.name)
			}
		})
	}
}
