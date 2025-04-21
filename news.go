package main

import (
	"io"
	"net/http"
	"strings"
)

func FetchNewsHeadlines() ([]string, error) {
	resp, err := http.Get("https://feeds.bbci.co.uk/news/rss.xml") // BBC RSS feed
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	raw := string(body)

	// Very rough headline extraction from RSS
	var headlines []string
	for _, line := range strings.Split(raw, "\n") {
		if strings.Contains(line, "<title>") && !strings.Contains(line, "BBC News") {
			// Extract the headline between <title> and </title>
			start := strings.Index(line, "<title>") + 7
			end := strings.Index(line, "</title>")
			if start != -1 && end != -1 {
				headline := line[start:end]
				// Remove <![CDATA[ and ]]>
				headline = strings.ReplaceAll(headline, "<![CDATA[", "")
				headline = strings.ReplaceAll(headline, "]]>", "")
				headlines = append(headlines, headline)
			}
		}
	}

	return headlines, nil
}
