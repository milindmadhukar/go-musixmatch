package gomusixmatch

import (
	"context"

	mxmParams "github.com/milindmadhukar/go-musixmatch/params"
)

// Get the lyrics for track based on title and artist
//
// Parameters:
//     QueryTrack  - The song title
//     QueryArtist - The song artist
func (client *Client) GetMatcherLyrics(ctx context.Context, params ...mxmParams.Param) (*Lyrics, error) {

	var lyricsData lyrics
	err := client.get(ctx, "matcher.lyrics.get", &lyricsData, params...)
	if err != nil {
		return nil, err
	}

	return &lyricsData.LyricsData, nil

}

// In some cases you already have some informations about the track title, artist name, album etc.

// A possible strategy to get the corresponding lyrics could be:
// search musixmatch's catalogue with a perfect match,
// maybe try using the fuzzy search,
// maybe try again using artist aliases, and so on.
//
// Parameters:
//     QueryTrack  - The song title
//     QueryArtist - The song artist
//     QueryAlbum  - The song album
func (client *Client) GetMatcherTrack(ctx context.Context, params ...mxmParams.Param) (*Track, error) {
	var trackData track
	err := client.get(ctx, "matcher.track.get", &trackData, params...)
	if err != nil {
		return nil, err
	}

	return &trackData.TrackData, nil

}

// func (client *Client) GetMatcherSubtitle(ctx context.Context, params ...musixmatchParams.Param) () {
//
// }
