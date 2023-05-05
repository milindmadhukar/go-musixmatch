package params_test

import (
	"net/url"
	"testing"
	"time"

	p "github.com/milindmadhukar/go-musixmatch/params"
)

func assert(expected, got interface{}, t *testing.T) {
	if expected != got {
		t.Errorf("Expected - %v, but got %v!", expected, got)
	}
}

func TestParams(t *testing.T) {
	var currTime = time.Now()
	var parsedUrl string
	var params []p.Param
	params = append(
		params,
		p.Country("IN"),
		p.TrackID(123),
		p.ArtistID(234),
		p.AlbumID(456),
		p.CommonTrackID(567),
		p.TrackMbID(678),
		p.ArtistMbID(890),
		p.AlbumMbID(2345),
		p.QueryTrack("test track"),
		p.QueryArtist("test artist"),
		p.QueryTrackArtist("test track artist"),
		p.QueryLyrics("test lyrics"),
		p.QueryWriter("test writer"),
		p.Query("test query"),
		p.TrackISRC("abc123"),
		p.SortByTrackRating("ASC"),
		p.SortByReleaseDate("DESC"),
		p.SortByArtistRating("ASC"),
		p.Page(10),
		p.PageSize(20),
		p.HasLyrics(true),
		p.IsInstrumental(false),
		p.HasSubtitle(true),
		p.FilterByArtistID(3456),
		p.FilterByArtistMbID(5678),
		p.FilterByMusicGenreID(34),
		p.FilterByLyricsLanguage("en"),
		p.FilterByMininiumReleaseDate(currTime),
		p.FilterByMaximumReleaseDate(currTime),
		p.FilterBySubtitleLength(time.Minute),
		p.FilterBySubtitleLengthMaxDeviation(time.Hour),
		p.ChartName("hot"),
	)

	p := p.Params{
		UrlParams: url.Values{},
	}
	for _, param := range params {
		param(&p)
		if p.Err != nil {
			t.Error(p.Err.Error())
		}
	}

	urlParams := p.UrlParams.Encode()

	if urlParams != "" {
		parsedUrl = "http:/xyz.xyz?xyz=xyz" + "&" + urlParams
	}
	parsedParams, err := url.ParseQuery(parsedUrl)

	if err != nil {
		t.Error(err.Error())
	}

	assert(parsedParams.Get("country"), "in", t)
	assert(parsedParams.Get("track_id"), "123", t)
	assert(parsedParams.Get("artist_id"), "234", t)
	assert(parsedParams.Get("album_id"), "456", t)
	assert(parsedParams.Get("commontrack_id"), "567", t)
	assert(parsedParams.Get("track_mbid"), "678", t)
	assert(parsedParams.Get("artist_mbid"), "890", t)
	assert(parsedParams.Get("album_mbid"), "2345", t)
	assert(parsedParams.Get("q_track"), "test track", t)
	assert(parsedParams.Get("q_artist"), "test artist", t)
	assert(parsedParams.Get("q_track_artist"), "test track artist", t)
	assert(parsedParams.Get("q_lyrics"), "test lyrics", t)
	assert(parsedParams.Get("q_writer"), "test writer", t)
	assert(parsedParams.Get("q"), "test query", t)
	assert(parsedParams.Get("track_isrc"), "abc123", t)
	assert(parsedParams.Get("s_track_rating"), "ASC", t)
	assert(parsedParams.Get("s_track_release_date"), "DESC", t)
	assert(parsedParams.Get("s_artist_rating"), "ASC", t)
	assert(parsedParams.Get("page"), "10", t)
	assert(parsedParams.Get("page_size"), "20", t)
	assert(parsedParams.Get("f_has_lyrics"), "1", t)
	assert(parsedParams.Get("f_is_instrumental"), "0", t)
	assert(parsedParams.Get("f_has_subtitle"), "1", t)
	assert(parsedParams.Get("f_artist_id"), "3456", t)
	assert(parsedParams.Get("f_artist_mbid"), "5678", t)
	assert(parsedParams.Get("f_music_genre_id"), "34", t)
	assert(parsedParams.Get("f_lyrics_language"), "en", t)
	assert(parsedParams.Get("f_track_release_group_first_release_date_min"), currTime.Format("20060102"), t)
	assert(parsedParams.Get("f_track_release_group_first_release_date_max"), currTime.Format("20060102"), t)
	assert(parsedParams.Get("f_subtitle_length"), "60", t)
	assert(parsedParams.Get("f_subtitle_length_max_deviation"), "3600", t)
	assert(parsedParams.Get("chart_name"), "hot", t)
}
