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
	ID   int    `json:"artist_id"`
	Name string `json:"artist_name"`
	/*
		When you create a world wide music related service you have to take into consideration to display the artist name in the user's local language. These translation are also used as aliases to improve the search results.

		Let's use PSY for this example.

		Western people know him as PSY but korean want to see the original name 싸이.

		Using the name translations provided by musixmatch's api you can show to every user the writing they expect to see.

		Furthermore, when you search for "psy gangnam style" or "싸이 gangnam style" with our search/match api you will still be able to find the song.
	*/
	Comment             string          `json:"artist_comment"`
	NameTranslationList json.RawMessage `json:"artist_name_translation_list"`
	// An artist comment is a short snippet of text which can be mainly used for disambiguation.
	// The artist country is the born country of the artist/group
	Country   string `json:"artist_country"`
	AliasList []struct {
		Alias string `json:"artist_alias,omitempty"`
	} `json:"artist_alias_list,omitempty"`
	Rating     int    `json:"artist_rating"`
	TwitterURL string `json:"artist_twitter_url,omitempty"`
	Credits    struct {
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
	ID                  int             `json:"track_id"`
	Name                string          `json:"track_name"`
	NameTranslationList json.RawMessage `json:"track_name_translation_list"`
	Rating              int             `json:"track_rating"`
	CommontrackID       int             `json:"commontrack_id"`
	Instrumental        int             `json:"instrumental"`
	Explicit            int             `json:"explicit"`
	HasLyrics           int             `json:"has_lyrics"`
	HasSubtitles        int             `json:"has_subtitles"`
	HasRichsync         int             `json:"has_richsync"`
	NumFavourite        int             `json:"num_favourite"`
	AlbumID             int             `json:"album_id"`
	AlbumName           string          `json:"album_name"`
	ArtistID            int             `json:"artist_id"`
	ArtistName          string          `json:"artist_name"`
	ShareURL            string          `json:"track_share_url"`
	EditURL             string          `json:"track_edit_url"`
	Restricted          int             `json:"restricted"`
	UpdatedTime         time.Time       `json:"updated_time"`
	PrimaryGenres       struct {
		MusicGenreList []struct {
			MusicGenreData MusicGenre `json:"music_genre"`
		} `json:"music_genre_list"`
	} `json:"primary_genres"`
}

type Lyrics struct {
	ID                int       `json:"lyrics_id"`
	Explicit          int       `json:"explicit"`
	Body              string    `json:"lyrics_body"`
	ScriptTrackingURL string    `json:"script_tracking_url"`
	PixelTrackingURL  string    `json:"pixel_tracking_url"`
	LyricsCopyright   string    `json:"lyrics_copyright"`
	UpdatedTime       time.Time `json:"updated_time"`
}

type Snippet struct {
	ID                int       `json:"snippet_id"`
	Language          string    `json:"snippet_language"`
	Restricted        int       `json:"restricted"`
	Instrumental      int       `json:"instrumental"`
	Body              string    `json:"snippet_body"`
	ScriptTrackingURL string    `json:"script_tracking_url"`
	PixelTrackingURL  string    `json:"pixel_tracking_url"`
	HTMLTrackingURL   string    `json:"html_tracking_url"`
	UpdatedTime       time.Time `json:"updated_time"`
}

type MusicGenre struct {
	ID           int    `json:"music_genre_id"`
	Name         string `json:"music_genre_name"`
	NameExtended string `json:"music_genre_name_extended"`
	Vanity       string `json:"music_genre_vanity"`
}

type Album struct {
	ID   int    `json:"album_id"`
	Mbid string `json:"album_mbid"`
	Name string `json:"album_name"`
	/*
		The artist rating is a score 0-100 identifying how popular is an artist in musixmatch.

		You can use this information to build charts, for suggestions, to sort search results.
	*/
	Rating int `json:"album_rating"`
	// The album official release date can be used to sort an artist's albums view starting by the most recent one.
	ReleaseDate string `json:"album_release_date"`
	ArtistID    int    `json:"artist_id"`
	ArtistName  string `json:"artist_name"`
	/*
	   For most of the albums we provide two groups of music genres. Primary and secondary. This information can be used to help user navigate albums by genre.

	   An example could be:
	       Primary genere: POP
	   	  Secondary genre: K-POP or Mandopop
	*/
	PrimaryGenres struct {
		MusicGenreList []struct {
			MusicGenreData MusicGenre `json:"music_genre"`
		} `json:"music_genre_list"`
	} `json:"primary_genres"`
	Pline       string    `json:"album_pline"`
	Copyright   string    `json:"album_copyright"`
	Label       string    `json:"album_label"`
	Restricted  int       `json:"restricted"`
	UpdatedTime time.Time `json:"updated_time"`
	ExternalIds struct {
		Spotify     []string `json:"spotify"`
		Itunes      []string `json:"itunes"`
		AmazonMusic []string `json:"amazon_music"`
	} `json:"external_ids"`
}

// HACK: I don't like the way JSON is being parsed by these not so useful structs. I also want to provide the end user withe request URL.

type album struct {
	AlbumData Album `json:"album"`
}

type albumList struct {
	AlbumList []struct {
		AlbumData Album `json:"album"`
	} `json:"album_list"`
}

type track struct {
	TrackData Track `json:"track"`
}

type trackList struct {
	TrackList []struct {
		TrackData Track `json:"track"`
	} `json:"track_list"`
}

type musicGenreList struct {
	MusicGenreList []struct {
		MusicGenreData MusicGenre `json:"music_genre"`
	} `json:"music_genre_list"`
}

type artist struct {
	ArtistData Artist `json:"artist"`
}

type artistList struct {
	ArtistList []struct {
		ArtistData Artist `json:"artist"`
	} `json:"artist_list"`
}

type lyrics struct {
	LyricsData Lyrics `json:"lyrics"`
}

type snippet struct {
	SnippetData Snippet `json:"snippet"`
}
