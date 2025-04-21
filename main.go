package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	InitDB()

	headlines, err := FetchNewsHeadlines()
	if err == nil {
		for _, h := range headlines[:3] { // limit to 3 for now
			prompt := fmt.Sprintf("Summarize this headline in one philosophical or reflective thought: \"%s\"", h)

			// Generate a thought based on the headline
			fmt.Println("📰 Generating thought for headline:", prompt)

			summary, err := GenerateThought(prompt)
			if err == nil {
				fmt.Println("📰 News thought:", summary)

				// Use your preferred news source name
				err = SaveNewsMemory("AutoNews", h, strings.TrimSpace(summary))
				if err != nil {
					fmt.Println("❌ Failed to save news memory:", err)
				}
			}

		}
	}

	for {
		// Pull memory + web memory
		recent := GetRecentThoughts(5)
		webmem := GetRecentWebMemories(3)
		news := GetRecentNewsMemories(3)

		// Combine them
		context := strings.Join(append(append(recent, webmem...), news...), "\n")

		//fmt.Println("🧠 Combined context:")
		//fmt.Println(context)

		prompt := fmt.Sprintf("Based on this combined context:\n%s\nWhat fetch next?", context)
		thought, err := GenerateThought(prompt)
		if err != nil {
			fmt.Println("Generation error:", err)
			continue
		}
		fmt.Println("🧠 Generated thought:", thought)

		thought = strings.TrimSpace(thought)
		// string is multiple lines, so we need to check each line
		lines := strings.Split(thought, "\n")
		for _, line := range lines {
			// Check if the line containts "fetch:"
			// if it does, we need to fetch the web data
			if strings.Contains(line, "fetch:") {
				// Extract the query from the line
				query := strings.TrimSpace(strings.TrimPrefix(line, "fetch:"))
				fmt.Println("🌐 Fetching web data for:", query)
				summary, err := FetchAndStoreWebData(query)
				if err != nil {
					fmt.Println("❌ Web fetch failed:", err)
					continue
				}

				fmt.Println("📥 Web summary:", summary)
				SaveThought("Web data on '" + query + "': " + summary)
			} else {
				SaveThought(thought)
			}
		}
		time.Sleep(5 * time.Second)
	}
}
