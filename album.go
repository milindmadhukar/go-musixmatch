package gomusixmatch

import (
	"context"
	"fmt"

	musixmatchParams "github.com/milindmadhukar/go-musixmatch/params"
)

type albumGet struct {
	Message struct {
		Header struct {
			StatusCode  int     `json:"status_code"`
			ExecuteTime float64 `json:"execute_time"`
		} `json:"header"`
		Body struct {
			AlbumData Album `json:"album"`
		} `json:"body"`
	} `json:"message"`
}

type albumTracksGet struct {
	Message struct {
		Header struct {
			StatusCode  int     `json:"status_code"`
			ExecuteTime float64 `json:"execute_time"`
			Available   int     `json:"available"`
		} `json:"header"`
		Body struct {
			TrackList []struct {
				TrackData Track `json:"track"`
			} `json:"track_list"`
		} `json:"body"`
	} `json:"message"`
}

func (client *Client) GetAlbum(ctx context.Context, params ...musixmatchParams.Param) (*albumGet, error) {
	url := fmt.Sprintf("%salbum.get?apikey=%s",
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

	var get_album albumGet

	err = client.get(ctx, url, &get_album)

	if err != nil {
		return nil, err
	}

	return &get_album, nil

}

func (client *Client) GetAlbumTracks(ctx context.Context, params ...musixmatchParams.Param) (*albumTracksGet, error) {
	url := fmt.Sprintf("%salbum.tracks.get?apikey=%s",
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

	var get_album_tracks albumTracksGet

	err = client.get(ctx, url, &get_album_tracks)

	if err != nil {
		return nil, err
	}

	return &get_album_tracks, nil

}
