package main

import (
	"context"
	"desume-client-go/desume"
	"fmt"
	"log"
)

func main() {
	ctx := context.Background()
	client := desume.NewClient()

	manga, err := client.GetMangaById(ctx, 1)
	mangas, err := client.GetMangas(ctx, nil)
	mangaChapter, err := client.GetMangaChapter(ctx, 1, 1)

	if err != nil {
		log.Fatalf("something went wrong: %+v", err)
	}

	fmt.Printf("Manga: %+v", manga)
	fmt.Printf("Mangas: %+v", mangas)
	fmt.Printf("MangaChapter: %+v", mangaChapter)
}
