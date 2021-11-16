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

// Get the lyrics for track based on title and artist
//
// Parameters:
//     QueryTrack  - The song title
//     QueryArtist - The song artist
func (client *Client) GetMatcherLyrics(ctx context.Context, params ...musixmatchParams.Param) (*matcherLyricsGet, error) {

	url := fmt.Sprintf("%smatcher.lyrics.get?apikey=%s",
		client.baseURL,
		client.apiKey)

	url, err := processParams(url, params...)
	if err != nil {
		return nil, err
	}

	var getLyrics matcherLyricsGet

	err = client.get(ctx, url, &getLyrics)

	if err != nil {
		return nil, err
	}

	getLyrics.Url = url

	return &getLyrics, nil

}

// In some cases you already have some informations about the track title, artist name, album etc.

// A possible strategy to get the corresponding lyrics could be:
// search our catalogue with a perfect match,
// maybe try using the fuzzy search,
// maybe try again using artist aliases, and so on.
//
// Parameters:
//     QueryTrack  - The song title
//     QueryArtist - The song artist
//     QueryAlbum  - The song album
func (client *Client) GetMatcherTrack(ctx context.Context, params ...musixmatchParams.Param) (*matcherTrackGet, error) {
	url := fmt.Sprintf("%smatcher.track.get?apikey=%s",
		client.baseURL,
		client.apiKey)

	url, err := processParams(url, params...)
	if err != nil {
		return nil, err
	}

	var getTrack matcherTrackGet

	err = client.get(ctx, url, &getTrack)

	if err != nil {
		return nil, err
	}

	getTrack.Url = url

	return &getTrack, nil

}

// func (client *Client) GetMatcherSubtitle(ctx context.Context, params ...musixmatchParams.Param) () {
//
// }
