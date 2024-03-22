package desume

const (
	KindManga    = "manga"
	KindManhwa   = "manhwa"
	KindManhua   = "manhua"
	KindOneShot  = "one_shot"
	KindComics   = "comics"
	OrderName    = "name"
	OrderPopular = "popular"
	OrderUpdated = "updated"
)

func GetDefaultMangaFilterParams() map[string]string {
	return map[string]string{
		"limit": "1",
		"kinds": KindManga,
		"Order": OrderPopular,
	}
}

func GetMangaFilterParams(
	page,
	limit,
	order,
	kinds,
	genres,
	search string,
) map[string]string {
	paramsWithKey := map[string]string{
		"page":   page,
		"limit":  limit,
		"search": search,
		"genres": genres,
		"kinds":  kinds,
		"order":  order,
	}

	currentParams := GetDefaultMangaFilterParams()

	for key, value := range paramsWithKey {
		if value != "" {
			currentParams[key] = value
		}
	}

	return currentParams
}
