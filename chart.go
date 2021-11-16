package gomusixmatch

import (
	"context"
	"fmt"

	musixmatchParams "github.com/milindmadhukar/go-musixmatch/params"
)

type topArtists struct {
	ArtistList []struct {
		ArtistData Artist `json:"artist"`
	} `json:"artist_list"`
	Url string `json:"url,omitempty"`
}

type topTracks struct {
	TrackList []struct {
		TrackData Track `json:"track"`
	} `json:"track_list"`
	Url string `json:"url,omitempty"`
}

// Gets the API response of the top artists of a given country.
//
// Parameters:
//     Country   - A valid country code.
//     Page      - Define the page number for paginated results.
//     PageSize  - Define the page size for paginated results. Range is 1 to 100.
func (client *Client) GetTopArtists(ctx context.Context, params ...musixmatchParams.Param) (*topArtists, error) {

	url := fmt.Sprintf("%schart.artists.get?apikey=%s",
		client.baseURL,
		client.apiKey)

	url, err := processParams(url, params...)
	if err != nil {
		return nil, err
	}

	var topArtists topArtists

	err = client.get(ctx, url, &topArtists)

	if err != nil {
		return nil, err
	}

	topArtists.Url = url

	return &topArtists, nil

}

// Gets the API response of the top songs of a given country.
//
// Parameters:
//     Country    - A valid 2 letters country code. Set XW as worldwide
//     Page       - Define the page number for paginated results.
//     PageSize  - Define the page size for paginated results. Range is 1 to 100.
//     ChartName - Select among available charts:
//         top : editorial chart.
//         hot : Most viewed lyrics in the last 2 hours.
//         mxmweekly : Most viewed lyrics in the last 7 days.
//         mxmweekly_new : Most viewed lyrics in the last 7 days limited to new releases only.
//     HasLyrics - When set to true, filters only contents with lyrics.
func (client *Client) GetTopTracks(ctx context.Context, params ...musixmatchParams.Param) (*topTracks, error) {

	url := fmt.Sprintf("%schart.tracks.get?apikey=%s",
		client.baseURL,
		client.apiKey)

	var topTracks topTracks

	url, err := processParams(url, params...)

	if err != nil {
		return nil, err
	}

	err = client.get(ctx, url, &topTracks)

	if err != nil {
		return nil, err
	}

	topTracks.Url = url

	return &topTracks, nil

}
