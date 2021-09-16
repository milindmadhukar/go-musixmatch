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

func (client *Client) GetArtist(ctx context.Context, params ...musixmatchParams.Param) (*artistGet, error) {
	url := fmt.Sprintf("%sartist.get?apikey=%s",
		client.baseURL,
		client.apiKey)

	url, err := processParams(url, params...)
	if err != nil {
		return nil, err
	}

	var get_artist artistGet

	err = client.get(ctx, url, &get_artist)

	if err != nil {
		return nil, err
	}

	get_artist.Url = url

	return &get_artist, nil

}

func (client *Client) SearchArtist(ctx context.Context, params ...musixmatchParams.Param) (*artistSearch, error) {

	url := fmt.Sprintf("%sartist.search?apikey=%s",
		client.baseURL,
		client.apiKey)

	url, err := processParams(url, params...)
	if err != nil {
		return nil, err
	}

	var search_artist artistSearch

	err = client.get(ctx, url, &search_artist)

	if err != nil {
		return nil, err
	}

	search_artist.Url = url

	return &search_artist, nil

}

func (client *Client) GetArtistAlbums(ctx context.Context, params ...musixmatchParams.Param) (*artistAlbumGet, error) {

	url := fmt.Sprintf("%sartist.albums.get?apikey=%s",
		client.baseURL,
		client.apiKey)

	url, err := processParams(url, params...)
	if err != nil {
		return nil, err
	}

	var get_artist_albums artistAlbumGet

	err = client.get(ctx, url, &get_artist_albums)

	if err != nil {
		return nil, err
	}

	get_artist_albums.Url = url
	return &get_artist_albums, nil

}

func (client *Client) GetRelatedArtists(ctx context.Context, params ...musixmatchParams.Param) (*artistRelatedGet, error) {
	url := fmt.Sprintf("%sartist.related.get?apikey=%s",
		client.baseURL,
		client.apiKey)

	url, err := processParams(url, params...)
	if err != nil {
		return nil, err
	}

	var get_related_artist artistRelatedGet

	err = client.get(ctx, url, &get_related_artist)

	if err != nil {
		return nil, err
	}

	get_related_artist.Url = url

	return &get_related_artist, nil

}
