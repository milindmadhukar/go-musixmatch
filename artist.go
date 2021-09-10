package gomusixmatch

import (
	"context"
	"fmt"

	musixmatchParams "github.com/milindmadhukar/go-musixmatch/params"
)

type artistGet struct {
	Message struct {
		Header struct {
			StatusCode  int     `json:"status_code"`
			ExecuteTime float64 `json:"execute_time"`
		} `json:"header"`
		Body struct {
			ArtistData Artist `json:"artist"`
		} `json:"body"`
	} `json:"message"`
}

type artistSearch struct {
	Message struct {
		Header struct {
			StatusCode  int     `json:"status_code"`
			ExecuteTime float64 `json:"execute_time"`
			Available   int     `json:"available"`
		} `json:"header"`
		Body struct {
			ArtistList []struct {
				ArtistData Artist `json:"artist"`
			} `json:"artist_list"`
		} `json:"body"`
	} `json:"message"`
}

type artistAlbumGet struct {
	Message struct {
		Header struct {
			StatusCode  int     `json:"status_code"`
			ExecuteTime float64 `json:"execute_time"`
			Available   int     `json:"available"`
		} `json:"header"`
		Body struct {
			AlbumList []struct {
				AlbumData Album `json:"album"`
			} `json:"album_list"`
		} `json:"body"`
	} `json:"message"`
}

type artistRelatedGet struct {
	Message struct {
		Header struct {
			StatusCode  int     `json:"status_code"`
			Available   int     `json:"available"`
			ExecuteTime float64 `json:"execute_time"`
		} `json:"header"`
		Body struct {
			ArtistList []struct {
				ArtistData Artist `json:"artist"`
			} `json:"artist_list"`
		} `json:"body"`
	} `json:"message"`
}

func (client *Client) GetArtist(ctx context.Context, params ...musixmatchParams.Param) (*artistGet, error) {
	url := fmt.Sprintf("%sartist.get?apikey=%s",
		client.baseURL,
		client.apiKey)

	options, err := processParams(params...)
	if err != nil {
		return nil, err
	}
	urlParams := options.UrlParams.Encode()

	if urlParams != "" {
		url += "&" + urlParams
	}

	var get_artist artistGet

	err = client.get(ctx, url, &get_artist)

	if err != nil {
		return nil, err
	}

	return &get_artist, nil

}

func (client *Client) SearchArtist(ctx context.Context, params ...musixmatchParams.Param) (*artistSearch, error) {

	url := fmt.Sprintf("%sartist.search?apikey=%s",
		client.baseURL,
		client.apiKey)

	options, err := processParams(params...)
	if err != nil {
		return nil, err
	}
	urlParams := options.UrlParams.Encode()

	if urlParams != "" {
		url += "&" + urlParams
	}

	var search_artist artistSearch

	err = client.get(ctx, url, &search_artist)

	if err != nil {
		return nil, err
	}

	return &search_artist, nil

}

func (client *Client) GetArtistAlbums(ctx context.Context, params ...musixmatchParams.Param) (*artistAlbumGet, error) {

	url := fmt.Sprintf("%sartist.albums.get?apikey=%s",
		client.baseURL,
		client.apiKey)

	options, err := processParams(params...)
	if err != nil {
		return nil, err
	}
	urlParams := options.UrlParams.Encode()

	if urlParams != "" {
		url += "&" + urlParams
	}

	var get_artist_albums artistAlbumGet

	err = client.get(ctx, url, &get_artist_albums)

	if err != nil {
		return nil, err
	}

	return &get_artist_albums, nil

}

func (client *Client) GetRelatedArtists(ctx context.Context, params ...musixmatchParams.Param) (*artistRelatedGet, error) {
	url := fmt.Sprintf("%sartist.related.get?apikey=%s",
		client.baseURL,
		client.apiKey)

	options, err := processParams(params...)
	if err != nil {
		return nil, err
	}
	urlParams := options.UrlParams.Encode()

	if urlParams != "" {
		url += "&" + urlParams
	}

	var get_related_artist artistRelatedGet

	err = client.get(ctx, url, &get_related_artist)

	if err != nil {
		return nil, err
	}

	return &get_related_artist, nil

}
