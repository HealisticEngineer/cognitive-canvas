# cognitive-canvas

cognitive-canvas is a project designed to expand on ollama retention by creating a thought loop that generates ideas and insights based on past thoughts and external information sources like news feeds.

## Objectives
- Create a canvas for the model to draw on information from the news and other sources.
- Generate thoughts and ideas based on that information, simulating a cognitive process.

## Features
- **Thought Loop**: Continuously generates new thoughts based on recent memory context.
- **News Integration**: Fetches headlines from external news sources (e.g., BBC RSS feed) to provide fresh context.
- **Memory Persistence**: Stores thoughts in a SQLite database for future reference.
- **Web Integration**: Fetches and summarizes web search results for additional context.

## Installation
To get started with cognitive-canvas, follow these steps:

```bash
git clone https://github.com/HealisticEngineer/cognitive-canvas.git
cd cognitive-canvas
go run .
```

## Usage
1. Run the application using `go run .`.
2. The program will:
   - Initialize a SQLite database (`ollama_memory.db`) to store thoughts.
   - Fetch recent thoughts from memory and display them as context.
   - Generate a new thought based on the context using the `ollama` command-line tool.
   - Save the generated thought back into the database.
3. The loop will repeat every 60 seconds.
4. If a thought contains `fetch: <topic>`, the system will fetch web data for the specified topic and store the results in the database.

## Prerequisites
- **Go**: Ensure you have Go 1.24.2 or later installed.
- **ollama CLI**: Install the `ollama` command-line tool and ensure it is accessible in your system's PATH.
- **SQLite**: The application uses SQLite for memory persistence.

## File Structure
```
cognitive-canvas/
├── go.mod               # Go module definition
├── go.sum               # Dependency checksums
├── LICENSE              # Project license
├── main.go              # Entry point of the application
├── memory.go            # Memory management (SQLite integration)
├── news.go              # Fetches news headlines
├── thought.go           # Thought generation logic
├── web.go               # Fetches and summarizes web search results
├── ollama_memory.db     # SQLite database (auto-created)
├── flow.drawio          # Diagram of the application's workflow
├── README.md            # Project documentation
```

## Contributing
Contributions are welcome! To contribute:
1. Fork the repository.
2. Create a new branch for your feature or bugfix.
3. Submit a pull request with a detailed description of your changes.

## License
This project is licensed under the BSD 3-Clause License. See the [LICENSE](LICENSE) file for details.

## Milestones
- [x] Create a cognitive-canvas code that will create a thought loop.
- [x] Create a cognitive-canvas for the model to draw on information from the news and other sources.
- [ ] Enhance thought generation with more advanced AI models and context integration.
- [ ] Add a web interface for visualizing thoughts and memory.

## Acknowledgments
- Inspired by the concept of cognitive loops and AI-driven creativity.
- Uses the `ollama` CLI for thought generation and SQLite for memory persistence.