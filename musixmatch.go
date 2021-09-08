package gomusixmatch

import (
	"context"
	"encoding/json"
	"net/http"
)

type Client struct {
	http    *http.Client
	baseURL string
	apiKey  string
}

func New(apiKey string, httpClient *http.Client) *Client {
	client := &Client{
		http:    httpClient,
		baseURL: "http://api.musixmatch.com/ws/1.1/",
		apiKey:  apiKey,
	}
	return client
}

func (client *Client) get(ctx context.Context, url string, response interface{}) error {

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	// fmt.Println(req.URL)

	resp, err := client.http.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent {
		return nil
	}

	err = json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		return err
	}

	return nil

}
