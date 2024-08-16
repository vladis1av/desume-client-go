package desume

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func decodeResponse(r io.Reader, v interface{}) error {
	// Create a buffer for reading the response
	body, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	// Create a structure for the error
	var errorResponse MangaError

	// Trying to decode the response into an error structure
	if err := json.Unmarshal(body, &errorResponse); err == nil && errorResponse.Error != "" {
		// If the error is not empty, then return it
		return fmt.Errorf(errorResponse.Error)
	}

	// If there is no error, then decode the response into the target structure
	if err := json.Unmarshal(body, v); err != nil {
		return err
	}

	return nil
}

// GetMangas gets a list of manga filtered by the specified parameters.
func (c *Client) GetMangas(ctx context.Context, params Params) (*MangasFilteredResponse, error) {
	resp, err := c.sendRequest(ctx, http.MethodGet, "", params, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response MangasFilteredResponse
	if err := decodeResponse(resp.Body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetMangaById gets information about the manga by its ID.
func (c *Client) GetMangaById(ctx context.Context, id int64) (*MangaInfoResponse, error) {
	resp, err := c.sendRequest(ctx, http.MethodGet, strconv.FormatInt(id, 10), nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response MangaInfoResponse
	if err := decodeResponse(resp.Body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetMangaChapter gets information about the manga chapter by the manga and chapter IDs.
func (c *Client) GetMangaChapter(ctx context.Context, mangaId, chapterId int64) (*MangaChapterResponse, error) {
	endpoint := fmt.Sprintf("%v/chapter/%v", mangaId, chapterId)

	resp, err := c.sendRequest(ctx, http.MethodGet, endpoint, nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response MangaChapterResponse
	if err := decodeResponse(resp.Body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
