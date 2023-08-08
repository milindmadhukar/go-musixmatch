package gomusixmatch

import (
	"context"

	mxmParams "github.com/milindmadhukar/go-musixmatch/params"
)

/*
Gets the API response of the top artists of a given country.

Parameters:

	Country   - A valid country code.
	Page      - Define the page number for paginated results.
	PageSize  - Define the page size for paginated results. Range is 1 to 100.
*/
func (client *Client) GetTopArtists(ctx context.Context, params ...mxmParams.Param) ([]*Artist, error) {

	var artistsData artistList

	err := client.get(ctx, "chart.artists.get", &artistsData, params...)

	if err != nil {
		return nil, err
	}

	var artists []*Artist

	for _, artist := range artistsData.ArtistList {
		artists = append(artists, &artist.ArtistData)
	}

	return artists, nil

}

/*
Gets the API response of the top songs of a given country.

Parameters:

	Country    - A valid 2 letters country code. Set XW as worldwide
	Page       - Define the page number for paginated results.
	PageSize  - Define the page size for paginated results. Range is 1 to 100.
	ChartName - Select among available charts:
	    top : editorial chart.
	    hot : Most viewed lyrics in the last 2 hours.
	    mxmweekly : Most viewed lyrics in the last 7 days.
	    mxmweekly_new : Most viewed lyrics in the last 7 days limited to new releases only.
	HasLyrics - When set, filters only contents with lyrics.
*/
func (client *Client) GetTopTracks(ctx context.Context, params ...mxmParams.Param) ([]*Track, error) {

	var tracksData trackList

	err := client.get(ctx, "chart.tracks.get", &tracksData, params...)

	if err != nil {
		return nil, err
	}

	var tracks []*Track

	for _, track := range tracksData.TrackList {
		tracks = append(tracks, &track.TrackData)
	}

	return tracks, nil

}
