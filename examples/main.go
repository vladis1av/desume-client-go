package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/vladis1av/desume-client-go/desume"
)

func main() {
	ctx := context.Background()

	// Create with optional settings
	client := desume.NewClient(
		desume.WithBaseURL("https://desu.me/manga/api/"), // default https://desu.win/manga/api/
		desume.WithDisableCompression(true),
		desume.WithIdleConnTimeout(60*time.Second),
		desume.WithTimeout(30*time.Second),
		desume.WithMaxIdleConns(50),
		desume.WithRateLimiter(3, 1),
		desume.WithHeaders(http.Header{
			"User-Agent":      []string{"My-App/1.0"},
			"X-Custom-Header": []string{"CustomValue"},
		}),
	)
	// Recommended limit to prevent API blocking: // 3 requests per second with the possibility of briefly exceeding 1 request

	// Later we add or change headers
	client.SetHeader("X-Custom-Header", "CustomValue")

	// Getting information about manga by ID
	manga, err := client.GetMangaById(ctx, 1)
	if err != nil {
		if err == desume.ErrRateLimitExceeded {
			log.Println("Rate limit exceeded, waiting before retry...")
			time.Sleep(time.Second)
			manga, err = client.GetMangaById(ctx, 1)
		}
		if err != nil {
			log.Fatalf("manga error: %+v", err)
		}
	}

	// Getting a list of manga with default parameters
	mangas, err := client.GetMangas(ctx, desume.GetDefaultMangaFilterParams())
	if err != nil {
		if err == desume.ErrRateLimitExceeded {
			log.Println("Rate limit exceeded, waiting before retry...")
			time.Sleep(time.Second)
			mangas, err = client.GetMangas(ctx, desume.GetDefaultMangaFilterParams())
		}
		if err != nil {
			log.Fatalf("mangas error: %+v", err)
		}
	}

	// Getting manga information manga with chapters
	mangaChapter, err := client.GetMangaChapter(ctx, 1, 1)
	if err != nil {
		if err == desume.ErrRateLimitExceeded {
			log.Println("Rate limit exceeded, waiting before retry...")
			time.Sleep(time.Second)
			mangaChapter, err = client.GetMangaChapter(ctx, 1, 1)
		}
		if err != nil {
			log.Fatalf("mangaChapter error: %+v", err)
		}
	}

	// Getting a filtered list of manga
	mangasFiltered, err := client.GetMangas(ctx, desume.GetMangaFilterParams("1", "5", desume.OrderByUpdated, desume.KindManga, "game", "bleach"))
	if err != nil {
		if err == desume.ErrRateLimitExceeded {
			log.Println("Rate limit exceeded, waiting before retry...")
			time.Sleep(time.Second)
			mangasFiltered, err = client.GetMangas(ctx, desume.GetMangaFilterParams("1", "5", desume.OrderByUpdated, desume.KindManga, "game", "bleach"))
		}
		if err != nil {
			log.Fatalf("mangasFiltered error: %+v", err)
		}
	}

	// Output of results
	fmt.Printf("Manga: %+v\n", manga.Response.ID)
	fmt.Printf("Mangas: %+v\n", mangas.Response[0].ID)
	fmt.Printf("MangaChapter: %+v\n", mangaChapter.Response.ID)
	fmt.Printf("MangasFiltered: %+v\n", mangasFiltered.Response[0].ID)
}
