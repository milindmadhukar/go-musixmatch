package gomusixmatch_test

import (
	"context"
	"net/http"
	"testing"

	musixmatch "github.com/milindmadhukar/go-musixmatch"
	"github.com/milindmadhukar/go-musixmatch/params"
)

var client = musixmatch.New("<YOUR API KEY>", http.DefaultClient)

// INFO: Martin Garrix Musixmatch ID : 24407895
// INFO: AREA21 Musixmatch ID : 50722792
// INFO: Greatest Hits Vol. 1 Album ID : 47919158
// INFO: Martin Garrix ft. Bonn - High On Life ID : 154554016

const martinID = 24407895
const a21ID = 50722792
const a21AlbumID = 47919158
const highOnLifeID = 154554016

func assert(expected, got interface{}, t *testing.T) {
	if expected != got {
		t.Errorf("Expected - %v, but got %v!", expected, got)
	}
}

func TestGetAlbum(t *testing.T) {

	album, err := client.GetAlbum(context.Background(), params.AlbumID(a21AlbumID))

	if err != nil {
		t.Error(err.Error())
	}

	assert(album.Name, "Greatest Hits Vol. 1", t)
	assert(album.ArtistName, "AREA21 feat. Martin Garrix & Maejor", t)
	assert(album.ArtistID, a21ID, t)
	assert(album.ID, a21AlbumID, t)
}

func TestGetAlbumTracks(t *testing.T) {

	album_tracks, err := client.GetAlbumTracks(context.Background(), params.AlbumID(a21AlbumID))

	if err != nil {
		t.Error(err.Error())
	}

	assert((*album_tracks)[0].Name, "21", t)
}

func TestGetArtist(t *testing.T) {

	artist, err := client.GetArtist(context.Background(), params.ArtistID(martinID))

	if err != nil {
		t.Error(err.Error())
	}
	assert(artist.Name, "Martin Garrix", t)

}

func TestSearchArtist(t *testing.T) {
	artists, err := client.SearchArtist(context.Background(), params.QueryArtist("Martin Garrix"))

	if err != nil {
		t.Error(err.Error())
	}

	assert((*artists)[0].Name, "Martin Garrix", t)
}

func TestGetArtistAlbums(t *testing.T) {
	albums, err := client.GetArtistAlbums(context.Background(), params.ArtistID(a21ID))
	if err != nil {
		t.Error(err.Error())
	}

	assert((*albums)[0].Name, "Greatest Hits Vol. 1", t)
}

func TestGetMatcherLyrics(t *testing.T) {
	lyrics, err := client.GetMatcherLyrics(
		context.Background(),
		params.QueryTrack("High On Life"),
		params.QueryArtist("Martin Garrix"),
	)
	if err != nil {
		t.Error(err.Error())
	}

	assert(lyrics.Body[:28], "Killed the demons of my mind", t)
}

func TestGetMatcherTrack(t *testing.T) {
	track, err := client.GetMatcherTrack(
		context.Background(),
		params.QueryTrack("high on life"),
		params.QueryArtist("Martin Garrix"),
	)

	if err != nil {
		t.Error(err.Error())
	}

	assert(track.Name, "High On Life", t)

}

func TestGetTrack(t *testing.T) {
	track, err := client.GetTrack(
		context.Background(),
		params.TrackID(highOnLifeID),
	)

	if err != nil {
		t.Error(err.Error())
	}

	assert(track.Name, "High On Life", t)
}

func TestSearchTrack(t *testing.T) {
	tracks, err := client.SearchTrack(
		context.Background(),
		params.QueryTrack("high on life"),
		params.QueryArtist("martin garrix"),
	)

	if err != nil {
		t.Error(err.Error())
	}

	assert((*tracks)[0].Name, "High On Life", t)
}

func TestGetTrackLyrics(t *testing.T) {
	lyrics, err := client.GetTrackLyrics(
		context.Background(),
		params.TrackID(highOnLifeID),
	)

	if err != nil {
		t.Error(err.Error())
	}

	assert(lyrics.Body[:28], "Killed the demons of my mind", t)

}

func TestGetTrackSnippet(t *testing.T) {
	snippet, err := client.GetTrackSnippet(
		context.Background(),
		params.TrackID(highOnLifeID),
	)

	if err != nil {
		t.Error(err.Error())
	}

	assert(snippet.Body, "High on life 'til the day we die", t)

}

// TODO: TestGetRelatedArtists
// TODO: TestGetMusicGenres
