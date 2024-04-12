package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/vladis1av/desume-client-go/desume"
)

func main() {
	ctx := context.Background()

	client := desume.NewClient(desume.WithTimeout(30*time.Second), desume.WithMaxIdleConns(50))

	manga, err := client.GetMangaById(ctx, 1)
	if err != nil {
		log.Fatalf("manga error: %+v", err)
	}

	mangas, err := client.GetMangas(ctx, desume.GetDefaultMangaFilterParams())
	if err != nil {
		log.Fatalf("mangas error: %+v", err)
	}

	mangaChapter, err := client.GetMangaChapter(ctx, 1, 1)
	if err != nil {
		log.Fatalf("mangaChapter error: %+v", err)
	}

	mangasFiltered, err := client.GetMangas(ctx, desume.GetMangaFilterParams("1", "5", desume.OrderByUpdated, desume.KindManga, "game", "bleach"))
	if err != nil {
		log.Fatalf("mangasFiltered error: %+v", err)
	}

	fmt.Printf("Manga: %+v\n", manga.Response.ID)
	fmt.Printf("Mangas: %+v\n", mangas.Response[0].ID)
	fmt.Printf("MangaChapter: %+v\n", mangaChapter.Response.ID)
	fmt.Printf("mangasFiltered: %+v\n", mangasFiltered.Response[0].ID)
}
