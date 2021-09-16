package gomusixmatch

import (
	"encoding/json"
	"time"
)

type musixMatchResponse struct {
	Message struct {
		Header struct {
			StatusCode  int     `json:"status_code"`
			ExecuteTime float64 `json:"execute_time"`
			Available   int     `json:"available,omitempty"`
			Hint        string  `json:"hint,omitempty"`
			Confidence  int     `json:"confidence,omitempty"`
			Mode        string  `json:"mode,omitempty"`
			Cached      int     `json:"cached,omitempty"`
		} `json:"header"`
		Body json.RawMessage `json:"body"`
	} `json:"message"`
}

type Artist struct {
	ArtistID                  int           `json:"artist_id"`
	ArtistName                string        `json:"artist_name"`
	ArtistNameTranslationList []interface{} `json:"artist_name_translation_list"`
	ArtistComment             string        `json:"artist_comment"`
	ArtistCountry             string        `json:"artist_country"`
	ArtistAliasList           []struct {
		ArtistAlias string `json:"artist_alias,omitempty"`
	} `json:"artist_alias_list,omitempty"`
	ArtistRating     int    `json:"artist_rating"`
	ArtistTwitterURL string `json:"artist_twitter_url,omitempty"`
	ArtistCredits    struct {
		ArtistList []Artist `json:"artist_list,omitempty"`
	} `json:"artist_credits"`
	Restricted    int       `json:"restricted,omitempty"`
	UpdatedTime   time.Time `json:"updated_time,omitempty"`
	BeginDateYear string    `json:"begin_date_year,omitempty"`
	BeginDate     string    `json:"begin_date,omitempty"`
	EndDateYear   string    `json:"end_date_year,omitempty"`
	EndDate       string    `json:"end_date,omitempty"`
}

type Track struct {
	TrackID                  int           `json:"track_id"`
	TrackName                string        `json:"track_name"`
	TrackNameTranslationList []interface{} `json:"track_name_translation_list"`
	TrackRating              int           `json:"track_rating"`
	CommontrackID            int           `json:"commontrack_id"`
	Instrumental             int           `json:"instrumental"`
	Explicit                 int           `json:"explicit"`
	HasLyrics                int           `json:"has_lyrics"`
	HasSubtitles             int           `json:"has_subtitles"`
	HasRichsync              int           `json:"has_richsync"`
	NumFavourite             int           `json:"num_favourite"`
	AlbumID                  int           `json:"album_id"`
	AlbumName                string        `json:"album_name"`
	ArtistID                 int           `json:"artist_id"`
	ArtistName               string        `json:"artist_name"`
	TrackShareURL            string        `json:"track_share_url"`
	TrackEditURL             string        `json:"track_edit_url"`
	Restricted               int           `json:"restricted"`
	UpdatedTime              time.Time     `json:"updated_time"`
	PrimaryGenres            struct {
		MusicGenreList []struct {
			MusicGenreData MusicGenre `json:"music_genre"`
		} `json:"music_genre_list"`
	} `json:"primary_genres"`
}

type Lyrics struct {
	LyricsID          int       `json:"lyrics_id"`
	Explicit          int       `json:"explicit"`
	LyricsBody        string    `json:"lyrics_body"`
	ScriptTrackingURL string    `json:"script_tracking_url"`
	PixelTrackingURL  string    `json:"pixel_tracking_url"`
	LyricsCopyright   string    `json:"lyrics_copyright"`
	UpdatedTime       time.Time `json:"updated_time"`
}

type Snippet struct {
	SnippetID         int       `json:"snippet_id"`
	SnippetLanguage   string    `json:"snippet_language"`
	Restricted        int       `json:"restricted"`
	Instrumental      int       `json:"instrumental"`
	SnippetBody       string    `json:"snippet_body"`
	ScriptTrackingURL string    `json:"script_tracking_url"`
	PixelTrackingURL  string    `json:"pixel_tracking_url"`
	HTMLTrackingURL   string    `json:"html_tracking_url"`
	UpdatedTime       time.Time `json:"updated_time"`
}

type MusicGenre struct {
	MusicGenreID           int    `json:"music_genre_id"`
	MusicGenreParentID     int    `json:"music_genre_parent_id"`
	MusicGenreName         string `json:"music_genre_name"`
	MusicGenreNameExtended string `json:"music_genre_name_extended"`
	MusicGenreVanity       string `json:"music_genre_vanity"`
}

type Album struct {
	AlbumID          int    `json:"album_id"`
	AlbumMbid        string `json:"album_mbid"`
	AlbumName        string `json:"album_name"`
	AlbumRating      int    `json:"album_rating"`
	AlbumReleaseDate string `json:"album_release_date"`
	ArtistID         int    `json:"artist_id"`
	ArtistName       string `json:"artist_name"`
	PrimaryGenres    struct {
		MusicGenreList []struct {
			MusicGenreData MusicGenre `json:"music_genre"`
		} `json:"music_genre_list"`
	} `json:"primary_genres"`
	AlbumPline     string    `json:"album_pline"`
	AlbumCopyright string    `json:"album_copyright"`
	AlbumLabel     string    `json:"album_label"`
	Restricted     int       `json:"restricted"`
	UpdatedTime    time.Time `json:"updated_time"`
	ExternalIds    struct {
		Spotify     []string `json:"spotify"`
		Itunes      []string `json:"itunes"`
		AmazonMusic []string `json:"amazon_music"`
	} `json:"external_ids"`
}
