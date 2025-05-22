# URL Shortener

A simple URL shortener service written in Go that allows you to shorten long URLs and redirect users to the original URLs using short codes.

## Features

- **URL Shortening**: Convert long URLs into short, manageable codes
- **Duplicate Prevention**: Same URLs generate the same short code (using SHA-256 hashing)
- **Redirection**: Access shortened URLs through short codes
- **In-Memory Storage**: Fast access with map-based storage
- **JSON API**: RESTful endpoints with JSON responses
- **Random Code Generation**: 16-character alphanumeric codes for uniqueness

## Project Structure

```
.
├── main.go                 # Main server entry point
├── urlShortenerHelper.go   # HTTP handlers and routing logic
└── helper.go              # Utility functions (hashing, random string generation)
```

## API Endpoints

### POST /save

Shortens a URL and returns a unique code.

**Request:**

- Method: `POST`
- Content-Type: `application/x-www-form-urlencoded`
- Body: `url=https://example.com/very/long/url`

**Response:**

```json
{
  "code": "A1b2C3d4E5f6G7h8"
}
```

**Error Responses:**

- `405 Method Not Allowed` - Only POST requests are accepted
- `400 Bad Request` - Invalid form data or missing URL field

### GET /{code}

Redirects to the original URL associated with the given code.

**Request:**

- Method: `GET`
- Path: `/{code}` (e.g., `/A1b2C3d4E5f6G7h8`)

**Response:**

- `302 Found` - Redirects to the original URL
- `404 Not Found` - Code doesn't exist

## Getting Started

### Prerequisites

- Go 1.16 or higher

### Installation & Running

1. Clone or download the project files
2. Navigate to the project directory
3. Run the server:

```bash
go run .
```

The server will start on `http://127.0.0.1:8080`

### Usage Examples

**Shorten a URL:**

```bash
curl -X POST -d "url=https://www.google.com/search?q=golang" http://127.0.0.1:8080/save
```

Response:

```json
{
  "code": "mK9pL2nQ8rS5tU7v"
}
```

**Access the shortened URL:**

```bash
curl -L http://127.0.0.1:8080/mK9pL2nQ8rS5tU7v
```

This will redirect you to the original URL.

## How It Works

1. **URL Submission**: When a URL is submitted via POST to `/save`, the system:

   - Generates a SHA-256 hash of the URL
   - Checks if this URL has been shortened before
   - If new, generates a random 16-character code
   - If existing, returns the previously generated code
   - Stores the mapping in memory

2. **URL Resolution**: When accessing `/{code}`:

   - The system looks up the code in the URL mapping
   - If found, redirects the user to the original URL
   - If not found, returns a 404 error

3. **Duplicate Handling**: The system uses SHA-256 hashing to ensure that the same URL always gets the same short code, preventing unnecessary duplicates.

## Technical Details

- **Storage**: Uses in-memory Go maps for fast access (data is lost when server restarts)
- **Code Generation**: 16-character codes using letters (a-z, A-Z) provide 52^16 possible combinations
- **Hashing**: SHA-256 is used to detect duplicate URLs
- **HTTP Status Codes**: Proper status codes for different scenarios (200, 302, 400, 404, 405)
