Sure! Here's the content formatted and ready to paste

# Gopher Adventure Storybook (CYOA)

A simple interactive storytelling web server written in Go. Users can read a branching story (choose-your-own-adventure style) about a curious little gopher exploring the world. The story is served via HTTP and navigated through query parameters.

---

## Features

- Serve a choose-your-own-adventure story from a JSON file
- Navigate between story arcs using links
- Easy to extend with new story arcs and pages
- Cleanly templated HTML rendering using Go’s `text/template` package

---

## Project Structure

```
.
├── gopher.json         # Story data in JSON format
├── main.go             # Entry point of the server
├── handler.go          # Error handling utilities
├── story.go            # Story data structure
├── page.html           # HTML template for rendering stories
└── README.md           # This file
```

---

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) 1.16 or higher installed

### Run the project

```bash
go run main.go
```

By default, it uses `gopher.json` as the story file and starts a server at:

```
http://localhost:8080
```

To specify a different JSON file:

```bash
go run main.go -file=my-story.json
```

---

## How It Works

- The server listens for requests on `/`.
- You can navigate story arcs by passing a query parameter like `?story=intro` or `?story=denver`.
- Each arc is rendered using `page.html`, which uses Go templates.

Example URL:

```
http://localhost:8080/?story=denver
```

---

## JSON Format

Each story arc is a key in a map with this structure:

```json
{
  "intro": {
    "title": "The Little Blue Gopher",
    "story": ["Once upon a time...", "Where should we go?"],
    "options": [
      {
        "text": "Go to New York",
        "arc": "new-york"
      },
      {
        "text": "Go to Denver",
        "arc": "denver"
      }
    ]
  }
}
```

---

## Extending the Story

Just add more arcs to the `gopher.json` file using the same format. Each arc must have:

- A `title` (string)
- A `story` (array of paragraphs)
- `options` (array of choices with text and a destination arc)

---

## Development Notes

### Error Handling

All critical errors are passed to `HandleError()` which panics and is caught by a `recover()` block in `main()` for safe shutdown.

---

## License

This project is provided as-is for educational/demo purposes. Feel free to modify and use it however you like.

---

## Credits

Inspired by the classic ["Choose Your Own Adventure"](https://en.wikipedia.org/wiki/Choose_Your_Own_Adventure) books and gophers everywhere.

```

Let me know if you want to include:
- a sample `gopher.json` file,
- instructions for deploying with Docker or a cloud service,
- or badge/shield graphics (e.g., Go version, build status).
```
