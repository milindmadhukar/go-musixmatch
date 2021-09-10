package params

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type Param func(*Params)

type Params struct {
	UrlParams url.Values
	Err       error
}

// Localization

// The country code of the desired country.
func Country(country_code string) Param {

	var err error

	_, found := findInSlice(countryCodes, country_code)

	if !found {
		err = errors.New(fmt.Sprintf("%s is not a valid country code. You may use XW for worldwide charts.", country_code))
	}

	return func(param *Params) {
		param.UrlParams.Set("country", country_code)
		param.Err = err
	}
}

// Objects

// Musixmatch track id
func TrackID(track_id int) Param {
	return func(param *Params) {
		param.UrlParams.Set("track_id", strconv.Itoa(track_id))
	}
}

// Musixmatch artist id
func ArtistID(artist_id int) Param {
	return func(param *Params) {
		param.UrlParams.Set("artist_id", strconv.Itoa(artist_id))
	}
}

// Musixmatch album id
func AlbumID(album_id int) Param {
	return func(param *Params) {
		param.UrlParams.Set("album_id", strconv.Itoa(album_id))
	}
}

// Musixmatch commontrack id
func CommonTrackID(commontrack_id int) Param {
	return func(param *Params) {
		param.UrlParams.Set("commontrack_id", strconv.Itoa(commontrack_id))
	}
}

// musicbrainz recording or track id
func TrackMbID(track_mbid int) Param {
	return func(param *Params) {
		param.UrlParams.Set("track_mbid", strconv.Itoa(track_mbid))
	}
}

// musicbrainz artist id
func ArtistMbID(artist_mbid int) Param {
	return func(param *Params) {
		param.UrlParams.Set("artist_mbid", strconv.Itoa(artist_mbid))
	}
}

// musicbrainz release id
func AlbumMbID(albim_mbid int) Param {
	return func(param *Params) {
		param.UrlParams.Set("albim_mbid", strconv.Itoa(albim_mbid))
	}
}

// Querying

// search for a text string among song titles
func QueryTrack(q_track string) Param {
	return func(param *Params) {
		param.UrlParams.Set("q_track", q_track)
	}
}

// search for a text string among artist names
func QueryArtist(q_artist string) Param {
	return func(param *Params) {
		param.UrlParams.Set("q_artist", q_artist)
	}
}

// search for a text string among lyrics
func QueryLyrics(q_lyrics string) Param {
	return func(param *Params) {
		param.UrlParams.Set("q_lyrics", q_lyrics)
	}
}

// search for a text string among song titles,artist names and lyrics
func Query(q string) Param {
	return func(param *Params) {
		param.UrlParams.Set("q", q)
	}
}

//Sorting

// Sort the results by our popularity index for tracks, possible values are ASC | DESC
func SortByTrackRating(order string) Param {
	var err error

	if l_order := strings.ToLower(order); (l_order != "asc") && (l_order != "desc") {
		err = errors.New("Please provide a valid sorting method. Options are ASC or DESC")
	}

	return func(param *Params) {
		param.UrlParams.Set("s_track_rating", order)
		param.Err = err
	}
}

// Sort the results by track release date, possible values are ASC | DESC
func SortByReleaseDate(order string) Param {
	var err error

	if l_order := strings.ToLower(order); (l_order != "asc") && (l_order != "desc") {
		err = errors.New("Please provide a valid sorting method. Options are ASC or DESC")
	}

	return func(param *Params) {
		param.UrlParams.Set("s_track_release_date", order)
		param.Err = err
	}
}

// Sort the results by our popularity index for artists, possible values are ASC | DESC
func SortByArtistRating(order string) Param {
	var err error

	if l_order := strings.ToLower(order); (l_order != "asc") && (l_order != "desc") {
		err = errors.New("Please provide a valid sorting method. Options are ASC or DESC")
	}

	return func(param *Params) {
		param.UrlParams.Set("s_artist_rating", order)
		param.Err = err
	}
}

// Result Set Pagination

// Request specific result page.
func Page(page int) Param {
	return func(param *Params) {
		param.UrlParams.Set("page", strconv.Itoa(page))
	}
}

// Specify number of items per result page (range 1 to 100)
func PageSize(page_size int) Param {
	var err error
	if page_size < 1 || page_size > 100 {
		err = errors.New("Invalid page size. Page size should be between 1 and 100.")
	}
	return func(param *Params) {
		param.UrlParams.Set("page_size", strconv.Itoa(page_size))
		param.Err = err
	}
}

// Filtering

// Filter by objects with available lyrics
func HasLyrics(has_lyrics bool) Param {
	var has_lyrics_int int
	if has_lyrics {
		has_lyrics_int = 1
	}
	return func(param *Params) {
		param.UrlParams.Set("f_has_lyrics", strconv.Itoa(has_lyrics_int))
	}
}

// Filter instrumental songs
func IsInstrumental(is_instrumental bool) Param {
	var is_instrumental_int int
	if is_instrumental {
		is_instrumental_int = 1
	}

	return func(param *Params) {
		param.UrlParams.Set("f_is_instrumental", strconv.Itoa(is_instrumental_int))
	}
}

// Filter by objects with available subtitles
func HasSubtitle(has_subtitle bool) Param {
	var has_subtitle_int int
	if has_subtitle {
		has_subtitle_int = 1
	}

	return func(param *Params) {
		param.UrlParams.Set("f_has_subtitle", strconv.Itoa(has_subtitle_int))
	}
}

// TODO
// f_music_genre_id - Filter by objects with a specific music category
// f_subtitle_length - Filter subtitles by a given duration in seconds
// f_subtitle_length_max_deviation - Apply a deviation to a given subtitle duration (in seconds)
// f_lyrics_language - Filter the tracks by lyrics language
// f_artist_id - Filter by objects with a given Musixmatch artist_id
// f_artist_mbid - Filter by objects with a given musicbrainz artist id

// Other

// Filter objects from a particular chart.
// Select among available charts:
//     top           : editorial chart
//     hot           : Most viewed lyrics in the last 2 hours
//     mxmweekly     : Most viewed lyrics in the last 7 days
//     mxmweekly_new : Most viewed lyrics in the last 7 days limited to new releases only
func ChartName(chart_name string) Param {

	var err error

	_, ok := chart_names[chart_name]
	if !ok {
		err = errors.New(fmt.Sprintf("%s is not a valid chart name.", chart_name))
	}

	return func(param *Params) {
		param.UrlParams.Set("chart_name", chart_name)
		param.Err = err
	}
}

// Helper functions

func findInSlice(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
