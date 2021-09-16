package gomusixmatch

import (
	"context"
	"fmt"

	musixmatchParams "github.com/milindmadhukar/go-musixmatch/params"
)

type matcherLyricsGet struct {
	LyricsData Lyrics `json:"lyrics"`
	Url        string `json:"url,omitempty"`
}

type matcherTrackGet struct {
	TrackData Track  `json:"track"`
	Url       string `json:"url,omitempty"`
}

func (client *Client) GetMatcherLyrics(ctx context.Context, params ...musixmatchParams.Param) (*matcherLyricsGet, error) {

	url := fmt.Sprintf("%smatcher.lyrics.get?apikey=%s",
		client.baseURL,
		client.apiKey)

	url, err := processParams(url, params...)
	if err != nil {
		return nil, err
	}

	var get_lyrics matcherLyricsGet

	err = client.get(ctx, url, &get_lyrics)

	if err != nil {
		return nil, err
	}

	get_lyrics.Url = url

	return &get_lyrics, nil

}

func (client *Client) GetMatcherTrack(ctx context.Context, params ...musixmatchParams.Param) (*matcherTrackGet, error) {
	url := fmt.Sprintf("%smatcher.track.get?apikey=%s",
		client.baseURL,
		client.apiKey)

	url, err := processParams(url, params...)
	if err != nil {
		return nil, err
	}

	var get_track matcherTrackGet

	err = client.get(ctx, url, &get_track)

	if err != nil {
		return nil, err
	}

	get_track.Url = url

	return &get_track, nil

}

// func (client *Client) GetMatcherSubtitle(ctx context.Context, params ...musixmatchParams.Param) () {
//
// }
