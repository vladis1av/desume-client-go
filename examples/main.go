package main

import (
	"context"
	"fmt"
	"log"

	"github.com/vladis1av/desume-client-go/desume"
)

func main() {
	ctx := context.Background()
	client := desume.NewClient()

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

	mangasFiltered, err := client.GetMangas(ctx, desume.GetMangaFilterParams("1", "5", desume.OrderUpdated, desume.KindManga, "game", ""))
	if err != nil {
		log.Fatalf("mangasFiltered error: %+v", err)
	}

	fmt.Printf("Manga: %+v", manga)
	fmt.Printf("Mangas: %+v", mangas)
	fmt.Printf("MangaChapter: %+v", mangaChapter)
	fmt.Printf("mangasFiltered: %+v", mangasFiltered)
}
