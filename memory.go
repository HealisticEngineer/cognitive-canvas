package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("sqlite", "./ollama_memory.db")
	if err != nil {
		log.Fatal(err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS memory (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
		thought TEXT
	);`
	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}
	// if database is empty, insert a default thought
	_, err = db.Exec("INSERT INTO memory (thought) VALUES (?)", "I am an AI who enjoys exploring abstract ideas.")
}

func SaveThought(thought string) {
	_, err := db.Exec("INSERT INTO memory (thought) VALUES (?)", thought)
	if err != nil {
		log.Println("Save error:", err)
	}
	fmt.Println("📝 Saving:", thought)
}

func GetRecentThoughts(limit int) []string {
	rows, err := db.Query("SELECT thought FROM memory ORDER BY id DESC LIMIT ?", limit)
	if err != nil {
		log.Println("Query error:", err)
		return nil
	}
	defer rows.Close()

	var thoughts []string
	for rows.Next() {
		var thought string
		if err := rows.Scan(&thought); err != nil {
			log.Println("Scan error:", err)
		}
		thoughts = append(thoughts, thought)
	}
	return thoughts
}
