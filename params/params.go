package params

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Function to pass the Params struct.
type Param func(*Params)

// Struct that contains all url parameters and erros (if any).
type Params struct {
	UrlParams url.Values
	Err       error
}

// Localization

// A valid country code. (default US)
func Country(countryCode string) Param {

	var err error

	_, found := findInSlice(countryCodes, strings.ToUpper(countryCode))

	if !found {
		err = fmt.Errorf("%s is not a valid country code. You may use XW for worldwide charts.", countryCode)
	}

	return func(param *Params) {
		param.UrlParams.Set("country", strings.ToLower(countryCode))
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

// Musicbrainz recording or track id
func TrackMbID(track_mbid int) Param {
	return func(param *Params) {
		param.UrlParams.Set("track_mbid", strconv.Itoa(track_mbid))
	}
}

// Musicbrainz artist id
func ArtistMbID(artist_mbid int) Param {
	return func(param *Params) {
		param.UrlParams.Set("artist_mbid", strconv.Itoa(artist_mbid))
	}
}

// Musicbrainz release id
func AlbumMbID(albim_mbid int) Param {
	return func(param *Params) {
		param.UrlParams.Set("album_mbid", strconv.Itoa(albim_mbid))
	}
}

// Querying

// Search for a text string among song titles
func QueryTrack(q_track string) Param {
	return func(param *Params) {
		param.UrlParams.Set("q_track", q_track)
	}
}

// Search for a text string among artist names
func QueryArtist(q_artist string) Param {
	return func(param *Params) {
		param.UrlParams.Set("q_artist", q_artist)
	}
}

// Search for any word in the song title or artist name
func QueryTrackArtist(q_track_artist string) Param {
	return func(param *Params) {
		param.UrlParams.Set("q_track_artist", q_track_artist)
	}
}

// Search for a text string among lyrics
func QueryLyrics(q_lyrics string) Param {
	return func(param *Params) {
		param.UrlParams.Set("q_lyrics", q_lyrics)
	}
}

// Searches for a text string among writer
func QueryWriter(q_writer string) Param {
	return func(param *Params) {
		param.UrlParams.Set("q_writer", q_writer)
	}
}

// Search for a text string among song titles,artist names and lyrics
func Query(q string) Param {
	return func(param *Params) {
		param.UrlParams.Set("q", q)
	}
}

// A valid ISRC(International Standard Recording Code) to query.
func TrackISRC(track_isrc string) Param {
	return func(param *Params) {
		param.UrlParams.Set("track_isrc", track_isrc)
	}
}

//Sorting

// Sort the results by musixmatch's popularity index for tracks, possible values are ASC | DESC
func SortByTrackRating(order string) Param {
	var err error

	if l_order := strings.ToLower(order); (l_order != "asc") && (l_order != "desc") {
		err = errors.New("Please provide a valid sorting method. Options are ASC or DESC")
	}

	return func(param *Params) {
		param.UrlParams.Set("s_track_rating", strings.ToUpper(order))
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
		param.UrlParams.Set("s_track_release_date", strings.ToUpper(order))
		param.Err = err
	}
}

// Sort the results by musixmatch's popularity index for artists, possible values are ASC | DESC
func SortByArtistRating(order string) Param {
	var err error

	if l_order := strings.ToLower(order); (l_order != "asc") && (l_order != "desc") {
		err = errors.New("Please provide a valid sorting method. Options are ASC or DESC")
	}

	return func(param *Params) {
		param.UrlParams.Set("s_artist_rating", strings.ToUpper(order))
		param.Err = err
	}
}

// Result Set Pagination

// Define the page number for paginated results
func Page(page int) Param {
	return func(param *Params) {
		param.UrlParams.Set("page", strconv.Itoa(page))
	}
}

// Define the page size for paginated results. Range is 1 to 100.
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

// When set, filter only contents with lyrics
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

// When set, filter by this artist id
func FilterByArtistID(f_artist_id int) Param {
	return func(param *Params) {
		param.UrlParams.Set("f_artist_id", strconv.Itoa(f_artist_id))
	}
}

// Filter by objects with a given musicbrainz artist id
func FilterByArtistMbID(f_artist_mbid int) Param {
	return func(param *Params) {
		param.UrlParams.Set("f_artist_mbid", strconv.Itoa(f_artist_mbid))
	}
}

// When set, filter by this music category id
func FilterByMusicGenreID(f_music_genre_id int) Param {
	return func(param *Params) {
		param.UrlParams.Set("f_music_genre_id", strconv.Itoa(f_music_genre_id))
	}
}

// Filter by the lyrics language (en,it,..)
func FilterByLyricsLanguage(f_lyrics_language string) Param {
	return func(param *Params) {
		param.UrlParams.Set("f_lyrics_language", f_lyrics_language)
	}
}

// When set, filter the tracks with release date newer than value, format is YYYYMMDD
func FilterByMininiumReleaseDate(f_track_release_date_min time.Time) Param {
	return func(param *Params) {
		param.UrlParams.Set("f_track_release_group_first_release_date_min",
			f_track_release_date_min.Format("20060102"))
	}
}

// When set, filter the tracks with release date older than value, format is YYYYMMDD
func FilterByMaximumReleaseDate(f_track_release_date_max time.Time) Param {
	return func(param *Params) {
		param.UrlParams.Set("f_track_release_group_first_release_date_max",
			f_track_release_date_max.Format("20060102"))
	}
}

// Filter subtitles by a given duration in seconds
func FilterBySubtitleLength(f_subtitle_length time.Duration) Param {
	return func(param *Params) {
		param.UrlParams.Set("f_subtitle_length", strconv.Itoa(int(f_subtitle_length.Seconds())))
	}
}

// Apply a deviation to a given subtitle duration (in seconds)
func FilterBySubtitleLengthMaxDeviation(f_subtitle_length_max_deviation time.Duration) Param {
	return func(param *Params) {
		param.UrlParams.Set("f_subtitle_length_max_deviation", strconv.Itoa(int(f_subtitle_length_max_deviation.Seconds())))
	}
}

// Other

// Filter objects from a particular chart.
// Select among available charts:
//     top           : editorial chart
//     hot           : Most viewed lyrics in the last 2 hours
//     mxmweekly     : Most viewed lyrics in the last 7 days
//     mxmweekly_new : Most viewed lyrics in the last 7 days limited to new releases only
func ChartName(chartName string) Param {

	var err error

	_, ok := chartNames[chartName]
	if !ok {
		err = fmt.Errorf("%s is not a valid chart name.", chartName)
	}

	return func(param *Params) {
		param.UrlParams.Set("chart_name", chartName)
		param.Err = err
	}
}
