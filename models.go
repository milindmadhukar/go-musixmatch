package gomusixmatch

import "time"

type Artist struct {
	ArtistID        int           `json:"artist_id"`
	ArtistMbid      string        `json:"artist_mbid"`
	ArtistName      string        `json:"artist_name"`
	ArtistAliasList []interface{} `json:"artist_alias_list"`
	ArtistRating    int           `json:"artist_rating"`
	UpdatedTime     time.Time     `json:"updated_time"`
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
			MusicGenre struct {
				MusicGenreID           int    `json:"music_genre_id"`
				MusicGenreParentID     int    `json:"music_genre_parent_id"`
				MusicGenreName         string `json:"music_genre_name"`
				MusicGenreNameExtended string `json:"music_genre_name_extended"`
				MusicGenreVanity       string `json:"music_genre_vanity"`
			} `json:"music_genre"`
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
