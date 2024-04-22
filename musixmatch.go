package gomusixmatch

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	mxmParams "github.com/milindmadhukar/go-musixmatch/params"
)

// Musixmatch request client that holds the API key and the http client.
type Client struct {
	HttpClient *http.Client
	BaseURL    string
	ApiKey     string
}

// Creates a new Musixmatch client to make requests to the api.
func New(apiKey string, httpClient *http.Client) *Client {
	client := Client{
		HttpClient: httpClient,
		BaseURL:    "https://api.musixmatch.com/ws/1.1",
		ApiKey:     apiKey,
	}
	return &client
}

func (client *Client) get(ctx context.Context, endpoint string, response interface{}, params ...mxmParams.Param) error {

	url := fmt.Sprintf("%s/%s?apikey=%s",
		client.BaseURL,
		endpoint,
		client.ApiKey)


	url, err := processParams(url, params...)
	if err != nil {
		return err
	}

  // fmt.Println(url)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	resp, err := client.HttpClient.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	var result musixMatchResponse
	err = json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		return err
	}

	switch result.Message.Header.StatusCode {

	case http.StatusNoContent:
		return errors.New("No content.")

	case http.StatusBadRequest:
		return errors.New("Bad Request. Please check your parameters.")

	case http.StatusUnauthorized:
		return errors.New("Authentication failed, probably because of invalid/missing API key.")

	case http.StatusPaymentRequired:
		return errors.New("The usage limit has been reached, either you exceeded per day requests limits or your balance is insufficient.")

	case http.StatusForbidden:
		return errors.New("You are not authorized to perform this operation.")

	case http.StatusNotFound:
		return errors.New("The requested resource was not found.")

	case http.StatusMethodNotAllowed:
		return errors.New("The requested method was not found.")

	case http.StatusInternalServerError:
		return errors.New("Oops. Something were wrong.")

	case http.StatusServiceUnavailable:
		return errors.New("Musixmatch's system is a bit busy at the moment and your request can’t be satisfied.")
	}

	err = json.Unmarshal(result.Message.Body, &response)

	if err != nil {
		return err
	}

	return nil

}

// Didn't do coz premium features not available.
// TODO: track.lyrics.post
// TODO: track.lyrics.mood.get
// TODO: track.richsync.get
// TODO: track.lyrics.translation.get
// TODO: track.subtitle.translation.get
// TODO: tracking.url.get
// TODO: catalogue.dump.get
// TODO: work.post
