package gomusixmatch

import (
	"context"
	"fmt"

	musixmatchParams "github.com/milindmadhukar/go-musixmatch/params"
)

type matcherLyricsGet struct {
	Message struct {
		Header struct {
			StatusCode  int     `json:"status_code"`
			ExecuteTime float64 `json:"execute_time"`
		} `json:"header"`
		Body struct {
			LyricsData Lyrics `json:"lyrics"`
		} `json:"body"`
	} `json:"message"`
}

type matcherTrackGet struct {
	Message struct {
		Header struct {
			StatusCode  int     `json:"status_code"`
			ExecuteTime float64 `json:"execute_time"`
			Confidence  int     `json:"confidence"`
			Mode        string  `json:"mode"`
			Cached      int     `json:"cached"`
		} `json:"header"`
		Body struct {
			TrackData Track `json:"track"`
		} `json:"body"`
	} `json:"message"`
}

func (client *Client) GetMatcherLyrics(ctx context.Context, params ...musixmatchParams.Param) (*matcherLyricsGet, error) {

	url := fmt.Sprintf("%smatcher.lyrics.get?apikey=%s",
		client.baseURL,
		client.apiKey)

	options, err := processParams(params...)
	if err != nil {
		return nil, err
	}
	urlParams := options.UrlParams.Encode()

	if urlParams != "" {
		url += "&" + urlParams
	}

	var get_lyrics matcherLyricsGet

	err = client.get(ctx, url, &get_lyrics)

	if err != nil {
		return nil, err
	}

	return &get_lyrics, nil

}

func (client *Client) GetMatcherTrack(ctx context.Context, params ...musixmatchParams.Param) (*matcherTrackGet, error) {
	url := fmt.Sprintf("%smatcher.track.get?apikey=%s",
		client.baseURL,
		client.apiKey)

	options, err := processParams(params...)
	if err != nil {
		return nil, err
	}
	urlParams := options.UrlParams.Encode()

	if urlParams != "" {
		url += "&" + urlParams
	}

	var get_track matcherTrackGet

	err = client.get(ctx, url, &get_track)

	if err != nil {
		return nil, err
	}

	return &get_track, nil

}

// func (client *Client) GetMatcherSubtitle(ctx context.Context, params ...musixmatchParams.Param) () {
//
// }
