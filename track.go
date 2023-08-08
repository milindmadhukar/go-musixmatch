package gomusixmatch

import (
	"context"

	mxmParams "github.com/milindmadhukar/go-musixmatch/params"
)

/*
Search for track in Musixmatch's database.

Parameters:

		QueryTrack                  - The song title
		QueryArtist                 - The song artist
		QueryLyrics                 - Any word in the lyrics
		QueryTrackArtist            - Any word in the song title or artist name
		QueryWriter                 - Search among writers
		Query                       - Any word in the song title or artist name or lyrics
		FilterByArtistID            - When set, filter by this artist id
		FilterByMusicGenreID        - When set, filter by this music category id
	 	FilterByLyricsLanguage      - Filter by the lyrics language (en,it,..)
		HasLyrics                   - When set, filter only contents with lyrics
		FilterByMininiumReleaseDate - When set, filter the tracks with release date newer than value, format is YYYYMMDD
		FilterByMininiumReleaseDate - When set, filter the tracks with release date older than value, format is YYYYMMDD
		SortByArtistRating          - Sort by our popularity index for artists (asc|desc)
		SortByTrackRating           - Sort by our popularity index for tracks (asc|desc)

TODO: what is quorum_factor?

	quorum_factor                                - Search only a part of the given query string.Allowed range is (0.1 – 0.9)
	Page                                         - Define the page number for paginated results
	PageSize                                    - Define the page size for paginated results. Range is 1 to 100.
*/
func (client *Client) SearchTrack(ctx context.Context, params ...mxmParams.Param) ([]*Track, error) {

	var trackData trackList

	err := client.get(ctx, "track.search", &trackData, params...)

	if err != nil {
		return nil, err
	}

	var tracks []*Track

	for _, track := range trackData.TrackList {
		tracks = append(tracks, &track.TrackData)
	}

	return tracks, nil

}

/*
Get a track info from musixmatch's database: title, artist, isrc(s), instrumental flag.

Parameters:

	CommonTrackID - The Musixmatch commontrack id
	TrackISRC     - A valid ISRC identifier
*/
func (client *Client) GetTrack(ctx context.Context, params ...mxmParams.Param) (*Track, error) {
	var trackData track

	err := client.get(ctx, "track.get", &trackData, params...)

	if err != nil {
		return nil, err
	}

	return &trackData.TrackData, nil

}

/*
Get the lyrics of a track.

Parameters:

	TrackID       - The Musixmatch track id
	CommonTrackID - The Musixmatch commontrack id
*/
func (client *Client) GetTrackLyrics(ctx context.Context, params ...mxmParams.Param) (*Lyrics, error) {
	var lyricsData lyrics

	err := client.get(ctx, "track.lyrics.get", &lyricsData, params...)

	if err != nil {
		return nil, err
	}

	return &lyricsData.LyricsData, nil

}

// func (client *Client) GetTrackLyricsMood (ctx context.Context, params ...musixmatchParams.Param) () {
//
// }

/*
Get the snippet for a given track.
A lyrics snippet is a very short representation of a song lyrics. It’s usually twenty to a hundred characters long and it’s calculated extracting a sequence of words from the lyrics.

Parameters:

	TrackID - The musixmatch track id
*/
func (client *Client) GetTrackSnippet(ctx context.Context, params ...mxmParams.Param) (*Snippet, error) {

	var snippetData snippet

	err := client.get(ctx, "track.snippet.get", &snippetData, params...)

	if err != nil {
		return nil, err
	}

	return &snippetData.SnippetData, nil

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
