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
	_, err = db.Exec("INSERT INTO memory (thought) VALUES (?)", "I am interested in the world around me. I want to learn more about current events and news.")
	_, err = db.Exec("INSERT INTO memory (thought) VALUES (?)", "Note to self: I can request more information by writing 'fetch: <your topic here>'. The system will search the web and give me new memory, I should try it out by getting current events.")
	// memory table for web memory
	createWebTable := `
	CREATE TABLE IF NOT EXISTS web_memory (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
		query TEXT,
		response TEXT
	);`
	_, err = db.Exec(createWebTable)
	if err != nil {
		log.Fatal(err)
	}
	// memory table for news memory
	createNewsTable := `
	CREATE TABLE IF NOT EXISTS news_memory (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
		source TEXT,
		headline TEXT,
		summary TEXT
	);`
	_, err = db.Exec(createNewsTable)
	if err != nil {
		log.Fatal(err)
	}
}

func SaveThought(thought string) {
	_, err := db.Exec("INSERT INTO memory (thought) VALUES (?)", thought)
	if err != nil {
		log.Println("Save error:", err)
	}
}

func SaveNewsMemory(source, headline, summary string) error {
	_, err := db.Exec("INSERT INTO news_memory (source, headline, summary) VALUES (?, ?, ?)", source, headline, summary)
	return err
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

func GetRecentWebMemories(limit int) []string {
	rows, err := db.Query("SELECT query, response FROM web_memory ORDER BY id DESC LIMIT ?", limit)
	if err != nil {
		log.Println("Web memory query error:", err)
		return nil
	}
	defer rows.Close()

	var webmemories []string
	for rows.Next() {
		var query, response string
		if err := rows.Scan(&query, &response); err == nil {
			webmemories = append(webmemories, fmt.Sprintf("[Web] %s → %s", query, response))
		}
	}
	return webmemories
}

func GetRecentNewsMemories(limit int) []string {
	rows, err := db.Query("SELECT source, headline, summary FROM news_memory ORDER BY id DESC LIMIT ?", limit)
	if err != nil {
		log.Println("News memory query error:", err)
		return nil
	}
	defer rows.Close()

	var news []string
	for rows.Next() {
		var src, headline, summary string
		if err := rows.Scan(&src, &headline, &summary); err == nil {
			news = append(news, fmt.Sprintf("[News] (%s) %s → %s", src, headline, summary))
		}
	}
	return news
}
