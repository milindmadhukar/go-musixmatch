package gomusixmatch

import (
	"context"
	"fmt"

	musixmatchParams "github.com/milindmadhukar/go-musixmatch/params"
)

type trackSearch struct {
	Message struct {
		Header struct {
			StatusCode  int     `json:"status_code"`
			ExecuteTime float64 `json:"execute_time"`
			Available   int     `json:"available,omitempty"`
		} `json:"header"`
		Body struct {
			TrackList []struct {
				TrackData Track `json:"track,omitempty"`
			} `json:"track_list"`
		} `json:"body"`
	} `json:"message"`
}

type trackGet struct {
	Message struct {
		Header struct {
			StatusCode  int     `json:"status_code"`
			ExecuteTime float64 `json:"execute_time"`
			Available   int     `json:"available,omitempty"`
		} `json:"header"`
		Body struct {
			TrackData Track `json:"track,omitempty"`
		} `json:"body"`
	} `json:"message"`
}

type trackLyricsGet struct {
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

type trackSnippetGet struct {
	Message struct {
		Header struct {
			StatusCode  int     `json:"status_code"`
			ExecuteTime float64 `json:"execute_time"`
		} `json:"header"`
		Body struct {
			SnippetData Snippet `json:"snippet"`
		} `json:"body"`
	} `json:"message"`
}

// Search for track in Musixmatch's database.
// q_track                                      - The song title
// q_artist                                     - The song artist
// q_lyrics                                     - Any word in the lyrics
// q_track_artist                               - Any word in the song title or artist name
// q_writer                                     - Search among writers
// q                                            - Any word in the song title or artist name or lyrics
// f_artist_id                                  - When set, filter by this artist id
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

	options, err := processParams(params...)
	if err != nil {
		return nil, err
	}
	urlParams := options.UrlParams.Encode()

	if urlParams != "" {
		url += "&" + urlParams
	}

	var search_results trackSearch

	err = client.get(ctx, url, &search_results)

	if err != nil {
		return nil, err
	}

	return &search_results, nil

}

// TODO
// List function

func (client *Client) GetTrack(ctx context.Context, params ...musixmatchParams.Param) (*trackGet, error) {
	url := fmt.Sprintf("%strack.get?apikey=%s",
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

	var get_results trackGet

	err = client.get(ctx, url, &get_results)

	if err != nil {
		return nil, err
	}

	return &get_results, nil

}

func (client *Client) GetTrackLyrics(ctx context.Context, params ...musixmatchParams.Param) (*trackLyricsGet, error) {
	url := fmt.Sprintf("%strack.lyrics.get?apikey=%s",
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

	var get_lyrics trackLyricsGet

	err = client.get(ctx, url, &get_lyrics)

	if err != nil {
		return nil, err
	}

	return &get_lyrics, nil

}

// func (client *Client) GetTrackLyricsMood (ctx context.Context, params ...musixmatchParams.Param) () {
//
// }

func (client *Client) GetTrackSnippet(ctx context.Context, params ...musixmatchParams.Param) (*trackSnippetGet, error) {

	url := fmt.Sprintf("%strack.snippet.get?apikey=%s",
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

	var get_snippet trackSnippetGet

	err = client.get(ctx, url, &get_snippet)

	if err != nil {
		return nil, err
	}

	return &get_snippet, nil

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
