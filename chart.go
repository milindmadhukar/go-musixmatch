package gomusixmatch

import (
	"context"
	"errors"
	"fmt"
)

type TopArtists struct {
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

type TopTracks struct {
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

func (topArtists *TopArtists) List() *[]Artist {

	var artists []Artist

	for _, artist := range topArtists.Message.Body.ArtistList {
		artists = append(artists, artist.ArtistData)
	}

	return &artists
}

func (client *Client) GetTopArtists(ctx context.Context, country string, page uint, page_size uint) (*TopArtists, error) {

	if page_size < 1 || page_size > 100 {
		return nil, errors.New("Invalid page size. Page size should be between 1 and 100.")
	}

	_, found := findInSlice(country_codes, country)

	if !found {
		return nil, errors.New(fmt.Sprintf("%s is not a valid country code. You may use XW for worldwide charts.", country))
	}

	url := fmt.Sprintf("%schart.artists.get?apikey=%s&country=%s&page=%d&page_size=%d&format=json",
		client.baseURL,
		client.apiKey,
		country,
		page,
		page_size)

	var topArtists TopArtists

	err := client.get(ctx, url, &topArtists)

	if err != nil {
		return nil, err
	}

	return &topArtists, nil

}

func (client *Client) GetTopTracks(ctx context.Context, country string, page uint, page_size uint, chart_name string, has_lyrics bool) (*TopTracks, error) {

	_, ok := chart_names[chart_name]

	if !ok {
		// Help text if errored maybe
		return nil, errors.New(fmt.Sprintf("%s is not a valid chart name.", chart_name))
	}

	_, found := findInSlice(country_codes, country)

	if !found {
		return nil, errors.New(fmt.Sprintf("%s is not a valid country code. You may use XW for worldwide charts.", country))
	}

	var has_lyrics_int int8
	if has_lyrics {
		has_lyrics_int = 1
	}
	url := fmt.Sprintf("%schart.tracks.get?apikey=%s&country=%s&page=%d&page_size=%d&chart_name=%s&f_has_lyrics=%d",
		client.baseURL,
		client.apiKey,
		country,
		page,
		page_size,
		chart_name,
		has_lyrics_int)

	var topTracks TopTracks

	err := client.get(ctx, url, &topTracks)

	if err != nil {
		return nil, err
	}

	return &topTracks, nil

}

func (topTracks *TopTracks) List() *[]Track {

	var tracks []Track

	for _, track := range topTracks.Message.Body.TrackList {
		tracks = append(tracks, track.TrackData)
	}

	return &tracks
}
