package desume

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (c *Client) GetMangas(ctx context.Context, params Params) (*MangasFilteredResponse, error) {
	resp, err := c.doRequest(ctx, http.MethodGet, "", params, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response MangasFilteredResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetMangaById(ctx context.Context, id int) (*MangaInfoResponse, error) {
	resp, err := c.doRequest(ctx, http.MethodGet, "", nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response MangaInfoResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetMangaChapter(ctx context.Context, mangaId, chapterId int) (*MangaChapterResponse, error) {
	endpoint := fmt.Sprintf("%v/chapter/%v", mangaId, chapterId)

	resp, err := c.doRequest(ctx, http.MethodGet, endpoint, nil, nil)
	if err != nil {
		log.Print("top")
		return nil, err
	}
	log.Println(endpoint)
	log.Println(resp)
	defer resp.Body.Close()

	var response MangaChapterResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Print("bottom")
		return nil, err
	}

	return &response, nil
}
