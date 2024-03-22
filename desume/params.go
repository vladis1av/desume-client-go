package desume

// Kind params option for filters
const (
	KindManga   = "manga"
	KindManhwa  = "manhwa"
	KindManhua  = "manhua"
	KindOneShot = "one_shot"
	KindComics  = "comics"
)

// OrderBy params option for filters
const (
	OrderByName    = "name"
	OrderByPopular = "popular"
	OrderByUpdated = "updated"
)

// GetDefaultMangaFilterParams gets the default filtering params
func GetDefaultMangaFilterParams() map[string]string {
	return map[string]string{
		"limit": "1",
		"kinds": KindManga,
		"order": OrderByPopular,
	}
}

// GetMangaFilterParams creates a parameter map to filter the manga list.
// Returns a parameter map that can be used to filter the manga list.
// If any of the passed parameters is "", it will not be included in the resulting map.
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
