package gomusixmatch

import (
	"context"
	"fmt"

	musixmatchParams "github.com/milindmadhukar/go-musixmatch/params"
)

type trackSearch struct {
	TrackList []struct {
		TrackData Track `json:"track,omitempty"`
	} `json:"track_list"`
	Url string `json:"url,omitempty"`
}

type trackGet struct {
	TrackData Track  `json:"track,omitempty"`
	Url       string `json:"url,omitempty"`
}

type trackLyricsGet struct {
	LyricsData Lyrics `json:"lyrics"`
	Url        string `json:"url,omitempty"`
}

type trackSnippetGet struct {
	SnippetData Snippet `json:"snippet"`
	Url         string  `json:"url,omitempty"`
}

// Search for track in Musixmatch's database.
//
// Parameters:
//
// QueryTrack                                      - The song title
// QueryArtist                                    - The song artist
// QueryLyrics                                  - Any word in the lyrics
// QueryWriter                                     - Search among writers
// Query                                            - Any word in the song title or artist name or lyrics
// ArtistID                                  - When set, filter by this artist id
// f_music_genre_id                             - When set, filter by this music category id
// f_lyrics_language                            - Filter by the lyrics language (en,it,..)
// f_has_lyrics                                 - When set, filter only contents with lyrics
// f_track_release_group_first_release_date_min - When set, filter the tracks with release date newer than value, format is YYYYMMDD
// f_track_release_group_first_release_date_max - When set, filter the tracks with release date older than value, format is YYYYMMDD
// s_artist_rating                              - Sort by our popularity index for artists (asc|desc)
// s_track_rating                               - Sort by our popularity index for tracks (asc|desc)
// quorum_factor                                - Search only a part of the given query string.Allowed range is (0.1 – 0.9)
// page                                         - Define the page number for paginated results
// page_size                                    - Define the page size for paginated results. Range is 1 to 100.
func (client *Client) SearchTrack(ctx context.Context, params ...musixmatchParams.Param) (*trackSearch, error) {

	url := fmt.Sprintf("%strack.search?apikey=%s",
		client.baseURL,
		client.apiKey)

	url, err := processParams(url, params...)
	if err != nil {
		return nil, err
	}

	var searchResults trackSearch

	err = client.get(ctx, url, &searchResults)

	if err != nil {
		return nil, err
	}

	searchResults.Url = url

	return &searchResults, nil

}

func (client *Client) GetTrack(ctx context.Context, params ...musixmatchParams.Param) (*trackGet, error) {

	url := fmt.Sprintf("%strack.get?apikey=%s",
		client.baseURL,
		client.apiKey)
	url, err := processParams(url, params...)
	if err != nil {
		return nil, err
	}

	var getResults trackGet

	err = client.get(ctx, url, &getResults)

	if err != nil {
		return nil, err
	}

	getResults.Url = url

	return &getResults, nil

}

func (client *Client) GetTrackLyrics(ctx context.Context, params ...musixmatchParams.Param) (*trackLyricsGet, error) {
	url := fmt.Sprintf("%strack.lyrics.get?apikey=%s",
		client.baseURL,
		client.apiKey)

	url, err := processParams(url, params...)
	if err != nil {
		return nil, err
	}

	var getLyrics trackLyricsGet

	err = client.get(ctx, url, &getLyrics)

	if err != nil {
		return nil, err
	}

	getLyrics.Url = url

	return &getLyrics, nil

}

// func (client *Client) GetTrackLyricsMood (ctx context.Context, params ...musixmatchParams.Param) () {
//
// }

func (client *Client) GetTrackSnippet(ctx context.Context, params ...musixmatchParams.Param) (*trackSnippetGet, error) {

	url := fmt.Sprintf("%strack.snippet.get?apikey=%s",
		client.baseURL,
		client.apiKey)

	url, err := processParams(url, params...)
	if err != nil {
		return nil, err
	}

	var getSnippet trackSnippetGet

	err = client.get(ctx, url, &getSnippet)

	if err != nil {
		return nil, err
	}

	getSnippet.Url = url

	return &getSnippet, nil

}

// func (client *Client) GetTrackSubtitle(ctx context.Context, params ...musixmatchParams.Param) () {
//
// }

// func (client *Client) GetTrackRichSync(ctx context.Context, params ...musixmatchParams.Param) () {
//
// }

// func (client *Client) GetTrackLyricsTranslation(ctx context.Context, params ...musixmatchParams.Param) () {
//
// }

// func (client *Client) GetTrackSubtitleTranslation(ctx context.Context, params ...musixmatchParams.Param) () {
//
// }
