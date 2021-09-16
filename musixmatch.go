package gomusixmatch

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
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
		baseURL: "https://api.musixmatch.com/ws/1.1/",
		apiKey:  apiKey,
	}
	return client
}

func (client *Client) get(ctx context.Context, url string, response interface{}) error {

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	resp, err := client.http.Do(req)

	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	resp.Body.Close()

	resp.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	// type apiResponse struct {
	// 	Message struct {
	// 		Header struct {
	// 			StatusCode  int     `json:"status_code"`
	// 			ExecuteTime float64 `json:"execute_time"`
	// 			Available   int     `json:"available,omitempty"`
	// 			Hint        string  `json:"hint,omitempty"`
	// 		} `json:"header"`
	// 		Body interface{} `json:"body"`
	// 	} `json:"message"`
	// }

	var result musixMatchResponse
	err = json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

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
