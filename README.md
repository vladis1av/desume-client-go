<h1 align="center">DesuMe API Client for Go</h1>

<p align="center">
  <a href="https://pkg.go.dev/github.com/vladis1av/desume-client-go">
    <img src="https://pkg.go.dev/badge/github.com/vladis1av/desume-client-go.svg" alt="Go Reference">
  </a>
  <a href="https://goreportcard.com/report/github.com/vladis1av/desume-client-go">
    <img src="https://goreportcard.com/badge/github.com/vladis1av/desume-client-go" alt="Go Report Card">
  </a>
  <a href="https://github.com/vladis1av/desume-client-go/blob/main/LICENSE">
    <img src="https://img.shields.io/github/license/vladis1av/desume-client-go" alt="License">
  </a>
</p>

<p align="center">
  Go wrapper for the DesuMe Web API
</p>

## 🚀 Features

- 🔄 Easy-to-use wrapper for DesuMe API
- ⚙️ Configurable HTTP client with timeout and connection pool settings
- 🛡️ Built-in rate limiting to prevent API abuse
- 🌐 Support for all major API endpoints
- 🎛️ Customizable request parameters
- 📝 Ability to set and modify headers dynamically

## 📦 Installation

To install the library, simply run:

```bash
go get github.com/vladis1av/desume-client-go
```

## 🏁 Quick Start

### Creating a Client

First, import the package and create a new client:

```go
import (
    "github.com/vladis1av/desume-client-go/desume"
    "time"
)

// Create with optional settings
client := desume.NewClient(
  desume.WithBaseURL("https://desu.me/manga/api/"), // default https://desu.win/manga/api/
  desume.WithDisableCompression(true),
  desume.WithIdleConnTimeout(60*time.Second),
  desume.WithTimeout(30*time.Second),
  desume.WithMaxIdleConns(50),
  desume.WithRateLimiter(2, 1),
// Recommended limit to prevent API blocking: // 3 requests per second with the possibility of briefly exceeding 1 request
)
```

### Setting Headers

Use WithHeaders to set a set of headers when creating the client. These headers will be used in all requests and can be overridden by using SetHeader:

```go
client := desume.NewClient(
  desume.WithHeaders(http.Header{
    "User-Agent":      []string{"My-App/1.0"},
    "X-Custom-Header": []string{"CustomValue"},
  }),
)
```
You can dynamically change or add headers after the client is created using SetHeader:

```go
client.SetHeader("Authorization", "Bearer token")
```
### Getting Manga by ID

To fetch information about a specific manga:

```go
manga, err := client.GetMangaById(ctx, 1)
if err != nil {
    log.Fatalf("Failed to get manga: %v", err)
}
fmt.Printf("Manga: %+v\n", manga.Response.ID)
```

### Fetching a List of Manga

To get a list of manga with default parameters:

```go
mangas, err := client.GetMangas(ctx, desume.GetDefaultMangaFilterParams())
if err != nil {
    log.Fatalf("Failed to get mangas: %v", err)
}
fmt.Printf("First manga in list: %+v\n", mangas.Response[0].ID)
```

### Getting Manga with Chapters

To fetch manga information including chapters:

```go
mangaChapter, err := client.GetMangaChapter(ctx, 1, 1)
if err != nil {
    log.Fatalf("Failed to get manga chapter: %v", err)
}
fmt.Printf("Manga chapter: %+v\n", mangaChapter.Response.ID)
```

### Filtering Manga

To get a filtered list of manga:

```go
filteredMangas, err := client.GetMangas(ctx, desume.GetMangaFilterParams("1", "5", desume.OrderByUpdated, desume.KindManga, "game", "bleach"))
if err != nil {
    log.Fatalf("Failed to get filtered mangas: %v", err)
}
fmt.Printf("First filtered manga: %+v\n", filteredMangas.Response[0].ID)
```

## 🔒 Rate Limiting

The client includes a built-in rate limiter. You can handle rate limit errors like this:

```go
manga, err := client.GetMangaById(ctx, 1)
if err == desume.ErrRateLimitExceeded {
    log.Println("Rate limit exceeded, waiting before retry...")
    time.Sleep(time.Second)
    manga, err = client.GetMangaById(ctx, 1)
}
```

## 📚 Documentation

For more detailed information about the available methods and options, please refer to the [GoDoc documentation](https://pkg.go.dev/github.com/vladis1av/desume-client-go).

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
