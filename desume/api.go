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
	// Создаем буфер для чтения ответа
	body, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	// Создаем структуру для ошибки
	var errorResponse MangaError

	// Пытаемся декодировать ответ в структуру ошибки
	if err := json.Unmarshal(body, &errorResponse); err == nil && errorResponse.Error != "" {
		// Если ошибка не пустая, то возвращаем ее
		return fmt.Errorf(errorResponse.Error)
	}

	// Если ошибки нет, то декодируем ответ в целевую структуру
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
func (c *Client) GetMangaById(ctx context.Context, id int) (*MangaInfoResponse, error) {
	resp, err := c.sendRequest(ctx, http.MethodGet, strconv.Itoa(id), nil, nil)
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
func (c *Client) GetMangaChapter(ctx context.Context, mangaId, chapterId int) (*MangaChapterResponse, error) {
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
