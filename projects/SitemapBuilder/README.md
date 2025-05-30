Here's a comprehensive `README.md` for your Go web crawler project:

---

# 🕷️ Go Web Crawler

A simple concurrent web crawler built in Go that recursively fetches internal links from a given base URL up to a specified depth.

---

## 📦 Overview

This crawler:

- Starts from a base URL (e.g. `https://example.com`)
- Follows internal links (`href` values that are _not_ full URLs)
- Crawls links up to a specified depth level
- Uses goroutines and a `sync.WaitGroup` for concurrency
- Avoids visiting the same path more than once
- Parses HTML using the `golang.org/x/net/html` package

---

## 🚀 Usage

### 🛠️ Build

```bash
go build -o crawler
```

### ▶️ Run

```bash
./crawler -url https://example.com -deep 3
```

### 📌 Flags

| Flag    | Description                          | Default      |
| ------- | ------------------------------------ | ------------ |
| `-url`  | The base URL to start crawling from  | _(Required)_ |
| `-deep` | The depth level to crawl recursively | `10`         |

---

## 🔍 Example Output

```bash
Crawling https://example.com/ at level 1
Crawling https://example.com/about at level 2
Crawling https://example.com/contact at level 2
...
```

---

## 🧩 Project Structure

### `main.go`

- Parses CLI arguments
- Manages concurrency and shared state
- Kicks off the crawling with `fetchPaths()`

### `parser.go`

- Parses HTML using Go's `html` tokenizer
- Extracts `href` attributes from anchor (`<a>`) tags
- Filters and returns relative links for further crawling

---

## 🔧 Key Functions

### `fetchPaths(urls []string, level int)`

Recursively crawls each URL and parses new internal links, spawning goroutines while limiting recursion depth.

### `getUrlContentPage(url string)`

Fetches the HTML content of a given URL.

### `ParseLinks(content string)`

Parses HTML content to extract all anchor (`<a>`) tags.

### `filterUrls(urls []string)`

Filters out full (external) URLs, keeping only relative paths.

---

## 🔐 Thread-Safety

To prevent race conditions:

- A `sync.Mutex` is used to lock shared access to the `paths` map.
- A `sync.WaitGroup` is used to ensure the main routine waits for all crawling goroutines to finish.

---

## 📦 Dependencies

- [golang.org/x/net/html](https://pkg.go.dev/golang.org/x/net/html) for HTML parsing

Install it with:

```bash
go get golang.org/x/net/html
```
