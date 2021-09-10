package gomusixmatch

import (
	"context"
	"fmt"

	musixmatchParams "github.com/milindmadhukar/go-musixmatch/params"
)

type topArtists struct {
	Message struct {
		Header struct {
			StatusCode  int     `json:"status_code"`
			ExecuteTime float64 `json:"execute_time"`
		} `json:"header"`
		Body struct {
			ArtistList []struct {
				ArtistData Artist `json:"artist"`
			} `json:"artist_list"`
		} `json:"body"`
	} `json:"message"`
}

type topTracks struct {
	Message struct {
		Header struct {
			StatusCode  int     `json:"status_code"`
			ExecuteTime float64 `json:"execute_time"`
		} `json:"header"`
		Body struct {
			TrackList []struct {
				TrackData Track `json:"track"`
			} `json:"track_list"`
		} `json:"body"`
	} `json:"message"`
}

// Gets the API response of the top artists of a given country.
// ctx       : The context in which the request is made.
// country   : A valid country code.
//page      : Define the page number for paginated results.
// page_size : Define the page size for paginated results. Range is 1 to 100.
func (client *Client) GetTopArtists(ctx context.Context, params ...musixmatchParams.Param) (*topArtists, error) {

	url := fmt.Sprintf("%schart.artists.get?apikey=%s",
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

	var topArtists topArtists

	err = client.get(ctx, url, &topArtists)

	if err != nil {
		return nil, err
	}

	return &topArtists, nil

}

// Gets the list of top artists as a slice of Artist structs.
func (topArtists *topArtists) List() *[]Artist {

	var artists []Artist

	for _, artist := range topArtists.Message.Body.ArtistList {
		artists = append(artists, artist.ArtistData)
	}

	return &artists
}

// Gets the API response of the top songs of a given country.
// ctx        : The context in which the request is made.
// country    : A valid 2 letters country code. Set XW as worldwide
// page       : Define the page number for paginated results.
// page_size  : Define the page size for paginated results. Range is 1 to 100.
// chart_name : Select among available charts:
//     top - editorial chart.
//     hot - Most viewed lyrics in the last 2 hours.
//     mxmweekly - Most viewed lyrics in the last 7 days.
//     mxmweekly_new - Most viewed lyrics in the last 7 days limited to new releases only.
// has_lyrics : When set to true, filters only contents with lyrics.
func (client *Client) GetTopTracks(ctx context.Context, params ...musixmatchParams.Param) (*topTracks, error) {

	url := fmt.Sprintf("%schart.tracks.get?apikey=%s",
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

	var topTracks topTracks

	err = client.get(ctx, url, &topTracks)

	if err != nil {
		return nil, err
	}

	return &topTracks, nil

}

// Gets the list of top tracks as a slice of Track structs.
func (topTracks *topTracks) List() *[]Track {

	var tracks []Track

	for _, track := range topTracks.Message.Body.TrackList {
		tracks = append(tracks, track.TrackData)
	}

	return &tracks
}
