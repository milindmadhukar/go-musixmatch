package gomusixmatch

import (
	"context"
	"fmt"

	musixmatchParams "github.com/milindmadhukar/go-musixmatch/params"
)

type albumGet struct {
	AlbumData Album  `json:"album"`
	Url       string `json:"url,omitempty"`
}

type albumTracksGet struct {
	TrackList []struct {
		TrackData Track `json:"track"`
	} `json:"track_list"`
	Url string `json:"url,omitempty"`
}

func (client *Client) GetAlbum(ctx context.Context, params ...musixmatchParams.Param) (*albumGet, error) {
	url := fmt.Sprintf("%salbum.get?apikey=%s",
		client.baseURL,
		client.apiKey)

	url, err := processParams(url, params...)
	if err != nil {
		return nil, err
	}

	var get_album albumGet

	err = client.get(ctx, url, &get_album)

	if err != nil {
		return nil, err
	}

	get_album.Url = url

	return &get_album, nil

}

func (client *Client) GetAlbumTracks(ctx context.Context, params ...musixmatchParams.Param) (*albumTracksGet, error) {
	url := fmt.Sprintf("%salbum.tracks.get?apikey=%s",
		client.baseURL,
		client.apiKey)

	url, err := processParams(url, params...)
	if err != nil {
		return nil, err
	}

	var get_album_tracks albumTracksGet

	err = client.get(ctx, url, &get_album_tracks)

	if err != nil {
		return nil, err
	}

	get_album_tracks.Url = url

	return &get_album_tracks, nil

}
