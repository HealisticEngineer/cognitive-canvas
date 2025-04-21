package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func FetchAndStoreWebData(query string) (string, error) {
	// Simple web search endpoint (DuckDuckGo HTML as a placeholder)
	url := fmt.Sprintf("https://html.duckduckgo.com/html/?q=%s", query)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	raw := string(body)

	// Simple content filtering
	summaryPrompt := fmt.Sprintf("Summarize this HTML search result for '%s' in one paragraph:\n%s", query, raw[:4000])
	summary, err := GenerateThought(summaryPrompt)
	if err != nil {
		return "", err
	}

	// Store in DB
	_, err = db.Exec("INSERT INTO web_memory (query, response) VALUES (?, ?)", query, summary)
	if err != nil {
		return "", err
	}

	return summary, nil
}

func GetWebMemories(limit int) []string {
	rows, err := db.Query("SELECT query, response FROM web_memory ORDER BY id DESC LIMIT ?", limit)
	if err != nil {
		log.Println("Web memory query error:", err)
		return nil
	}
	defer rows.Close()

	var memories []string
	for rows.Next() {
		var q, r string
		if err := rows.Scan(&q, &r); err != nil {
			log.Println("Scan error:", err)
		}
		memories = append(memories, fmt.Sprintf("Q: %s\nA: %s", q, r))
	}
	return memories
}
