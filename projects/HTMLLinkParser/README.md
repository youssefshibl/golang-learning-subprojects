# HTML Link Extractor

A simple Go command-line tool that parses HTML files and extracts all anchor (`<a>`) tags along with their text content and href attributes.

## Features

- Parses HTML files using Go's `golang.org/x/net/html` package
- Extracts all anchor tags from HTML documents
- Retrieves both the link text and href attribute for each link
- Uses goroutines and channels for concurrent processing
- Supports custom HTML file input via command-line flags

## Installation

Make sure you have Go installed on your system, then install the required dependency:

```bash
go mod init html-link-extractor
go get golang.org/x/net/html
```

## Usage

### Basic Usage

By default, the program looks for a file named `ex1.html` in the current directory:

```bash
go run main.go
```

### Custom HTML File

Specify a different HTML file using the `-file` flag:

```bash
go run main.go -file mypage.html
go run main.go -file /path/to/your/file.html
```

### Build and Run

You can also build the executable first:

```bash
go build -o link-extractor main.go
./link-extractor -file example.html
```

## Output

The program outputs each link found in the HTML file in the following format:

```
{Text: "Link Text" Href: "https://example.com"}
{Text: "Another Link" Href: "/relative/path"}
```

## Example

Given an HTML file with the following content:

```html
<!DOCTYPE html>
<html>
  <head>
    <title>Sample Page</title>
  </head>
  <body>
    <a href="https://google.com">Google</a>
    <a href="/about">About Us</a>
    <a href="mailto:contact@example.com">Contact</a>
  </body>
</html>
```

The program would output:

```
{Text: "Google" Href: "https://google.com"}
{Text: "About Us" Href: "/about"}
{Text: "Contact" Href: "mailto:contact@example.com"}
```

## How It Works

1. **Command-line Parsing**: Uses Go's `flag` package to accept an optional HTML file parameter
2. **HTML Parsing**: Opens and parses the HTML file using `golang.org/x/net/html`
3. **Node Traversal**: Recursively traverses the HTML DOM tree using goroutines
4. **Link Extraction**: Identifies anchor (`<a>`) elements and extracts:
   - **Text**: All text content within the anchor tag (including nested elements)
   - **Href**: The value of the `href` attribute
5. **Channel Communication**: Uses Go channels to communicate found links between goroutines

## Code Structure

- `main()`: Entry point, handles command-line arguments and coordinates the extraction process
- `getTargetNode()`: Recursively traverses HTML nodes to find anchor tags, runs as a goroutine
- `extractHref()`: Extracts the `href` attribute value from an anchor node
- `extractText()`: Recursively extracts all text content from an anchor node and its children
- `check()`: Simple error handling helper function

## Requirements

- Go 1.16 or higher
- `golang.org/x/net/html` package
