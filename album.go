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

// Get an album from Musixmatch's database: name, release_date, release_type, cover art.
//
// Parameters :
//     AlbumID - The musixmatch album id
func (client *Client) GetAlbum(ctx context.Context, params ...musixmatchParams.Param) (*albumGet, error) {
	url := fmt.Sprintf("%salbum.get?apikey=%s",
		client.baseURL,
		client.apiKey)

	url, err := processParams(url, params...)
	if err != nil {
		return nil, err
	}

	var getAlbum albumGet

	err = client.get(ctx, url, &getAlbum)

	if err != nil {
		return nil, err
	}

	getAlbum.Url = url

	return &getAlbum, nil

}

// Get the list of the songs of an album.
//
// Parameters:
//     AlbumID   - Musixmatch album id
//     AlbumMbId - Musicbrainz album id
//     HasLyrics - When set, filter only contents with lyrics
//     Page      - Define the page number for paginated results
//     PageSize  - Define the page size for paginated results. Range is 1 to 100.
func (client *Client) GetAlbumTracks(ctx context.Context, params ...musixmatchParams.Param) (*albumTracksGet, error) {
	url := fmt.Sprintf("%salbum.tracks.get?apikey=%s",
		client.baseURL,
		client.apiKey)

	url, err := processParams(url, params...)
	if err != nil {
		return nil, err
	}

	var getAlbumTracks albumTracksGet

	err = client.get(ctx, url, &getAlbumTracks)

	if err != nil {
		return nil, err
	}

	getAlbumTracks.Url = url

	return &getAlbumTracks, nil

}
