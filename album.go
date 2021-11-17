package gomusixmatch

import (
	"context"

	mxmParams "github.com/milindmadhukar/go-musixmatch/params"
)

/*Get an album from Musixmatch's database: name, release_date, release_type, cover art.

Parameters :
    AlbumID - The musixmatch album id
*/
func (client *Client) GetAlbum(ctx context.Context, params ...mxmParams.Param) (*Album, error) {
	var albumData album

	err := client.get(ctx, "album.get", &albumData, params...)

	if err != nil {
		return nil, err
	}

	return &albumData.AlbumData, nil

}

/*Get the list of the songs of an album.

Parameters:
    AlbumID   - Musixmatch album id
    AlbumMbId - Musicbrainz album id
    HasLyrics - When set, filter only contents with lyrics
    Page      - Define the page number for paginated results
    PageSize  - Define the page size for paginated results. Range is 1 to 100.
*/
func (client *Client) GetAlbumTracks(ctx context.Context, params ...mxmParams.Param) (*[]Track, error) {
	var tracksData trackList

	err := client.get(ctx, "album.tracks.get", &tracksData, params...)

	if err != nil {
		return nil, err
	}

	var tracks []Track

	for _, track := range tracksData.TrackList {
		tracks = append(tracks, track.TrackData)
	}

	return &tracks, nil

}
