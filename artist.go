package gomusixmatch

import (
	"context"
	"fmt"

	musixmatchParams "github.com/milindmadhukar/go-musixmatch/params"
)

type artistGet struct {
	ArtistData Artist `json:"artist"`
	Url        string `json:"url,omitempty"`
}

type artistSearch struct {
	ArtistList []struct {
		ArtistData Artist `json:"artist"`
	} `json:"artist_list"`
	Url string `json:"url,omitempty"`
}

type artistAlbumGet struct {
	AlbumList []struct {
		AlbumData Album `json:"album"`
	} `json:"album_list"`
	Url string `json:"url,omitempty"`
}

type artistRelatedGet struct {
	ArtistList []struct {
		ArtistData Artist `json:"artist"`
	} `json:"artist_list"`
	Url string `json:"url,omitempty"`
}

// Get the artist data from Musixmatch's database.
//
// Parameters:
//     ArtistID   - Musixmatch artist id
//     ArtistMbID - Musicbrainz artist id
func (client *Client) GetArtist(ctx context.Context, params ...musixmatchParams.Param) (*artistGet, error) {
	url := fmt.Sprintf("%sartist.get?apikey=%s",
		client.baseURL,
		client.apiKey)

	url, err := processParams(url, params...)
	if err != nil {
		return nil, err
	}

	var getArtist artistGet

	err = client.get(ctx, url, &getArtist)

	if err != nil {
		return nil, err
	}

	getArtist.Url = url

	return &getArtist, nil

}

// Search for artists in Musixmatch's database.
//
// Parameters:
//     QueryArtist - The song artist
//     ArtistID    - When set, filter by this artist id
//     ArtistMbID  - When set, filter by this artist musicbrainz id
//     Page        - Define the page number for paginated results
//     PageSize    - Define the page size for paginated results. Range is 1 to 100.
func (client *Client) SearchArtist(ctx context.Context, params ...musixmatchParams.Param) (*artistSearch, error) {

	url := fmt.Sprintf("%sartist.search?apikey=%s",
		client.baseURL,
		client.apiKey)

	url, err := processParams(url, params...)
	if err != nil {
		return nil, err
	}

	var searchArtist artistSearch

	err = client.get(ctx, url, &searchArtist)

	if err != nil {
		return nil, err
	}

	searchArtist.Url = url

	return &searchArtist, nil

}

// Get the album discography of an artist
//
// Parameters:
//     ArtistID          - Musixmatch artist id
//     ArtistMbID        - Musicbrainz artist id
//     GroupByAlbumName  - Group by Album Name
//     SortByReleaseDate - Sort by release date (asc|desc)
//     Page              - Define the page number for paginated results
//     PageSize          - Define the page size for paginated results. Range is 1 to 100.
func (client *Client) GetArtistAlbums(ctx context.Context, params ...musixmatchParams.Param) (*artistAlbumGet, error) {

	url := fmt.Sprintf("%sartist.albums.get?apikey=%s",
		client.baseURL,
		client.apiKey)

	url, err := processParams(url, params...)
	if err != nil {
		return nil, err
	}

	var getArtistAlbums artistAlbumGet

	err = client.get(ctx, url, &getArtistAlbums)

	if err != nil {
		return nil, err
	}

	getArtistAlbums.Url = url
	return &getArtistAlbums, nil

}

//Get a list of artists somehow related to a given one.
//
// Parameters:
//     ArtistID          - Musixmatch artist id
//     ArtistMbID        - Musicbrainz artist id
//     Page              - Define the page number for paginated results
//     PageSize          - Define the page size for paginated results. Range is 1 to 100.
func (client *Client) GetRelatedArtists(ctx context.Context, params ...musixmatchParams.Param) (*artistRelatedGet, error) {
	url := fmt.Sprintf("%sartist.related.get?apikey=%s",
		client.baseURL,
		client.apiKey)

	url, err := processParams(url, params...)
	if err != nil {
		return nil, err
	}

	var getRelatedArtist artistRelatedGet

	err = client.get(ctx, url, &getRelatedArtist)

	if err != nil {
		return nil, err
	}

	getRelatedArtist.Url = url

	return &getRelatedArtist, nil

}
