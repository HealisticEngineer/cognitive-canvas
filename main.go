package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	InitDB()

	for {
		recent := GetRecentThoughts(5)
		fmt.Println("🧠 Memory Context:")
		for _, t := range recent {
			fmt.Println(" -", t)
		}
		context := strings.Join(recent, "\n")

		prompt := fmt.Sprintf("Based on these past thoughts:\n%s\nWhat should I think about next?", context)
		thought, err := GenerateThought(prompt)
		if err != nil {
			fmt.Println("Generation error:", err)
			continue
		}

		fmt.Println("💭", strings.TrimSpace(thought))
		SaveThought(strings.TrimSpace(thought))

		time.Sleep(60 * time.Second)
	}
}
