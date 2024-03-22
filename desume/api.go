package desume

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// GetMangas gets a list of manga filtered by the specified parameters.
func (c *Client) GetMangas(ctx context.Context, params Params) (*MangasFilteredResponse, error) {
	resp, err := c.sendRequest(ctx, http.MethodGet, "", params, nil)
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

// GetMangaById gets information about the manga by its ID.
func (c *Client) GetMangaById(ctx context.Context, id int) (*MangaInfoResponse, error) {
	resp, err := c.sendRequest(ctx, http.MethodGet, strconv.Itoa(id), nil, nil)
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

// GetMangaChapter gets information about the manga chapter by the manga and chapter IDs.
func (c *Client) GetMangaChapter(ctx context.Context, mangaId, chapterId int) (*MangaChapterResponse, error) {
	endpoint := fmt.Sprintf("%v/chapter/%v", mangaId, chapterId)

	resp, err := c.sendRequest(ctx, http.MethodGet, endpoint, nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response MangaChapterResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
