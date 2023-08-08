package gomusixmatch

import (
	"context"

	mxmParams "github.com/milindmadhukar/go-musixmatch/params"
)

/*
Get the artist data from Musixmatch's database.

Parameters:

	ArtistID   - Musixmatch artist id
	ArtistMbID - Musicbrainz artist id
*/
func (client *Client) GetArtist(ctx context.Context, params ...mxmParams.Param) (*Artist, error) {
	var artistData artist

	err := client.get(ctx, "artist.get", &artistData, params...)

	if err != nil {
		return nil, err
	}

	return &artistData.ArtistData, nil

}

/*
Search for artists in Musixmatch's database.

Parameters:

	QueryArtist - The song artist
	ArtistID    - When set, filter by this artist id
	ArtistMbID  - When set, filter by this artist musicbrainz id
	Page        - Define the page number for paginated results
	PageSize    - Define the page size for paginated results. Range is 1 to 100.
*/
func (client *Client) SearchArtist(ctx context.Context, params ...mxmParams.Param) ([]*Artist, error) {

	var artistsData artistList

	err := client.get(ctx, "artist.search", &artistsData, params...)

	if err != nil {
		return nil, err
	}

	var artists []*Artist

	for _, artist := range artistsData.ArtistList {
		artists = append(artists, &artist.ArtistData)
	}

	return artists, nil

}

/*
Get the album discography of an artist

Parameters:

	ArtistID          - Musixmatch artist id
	ArtistMbID        - Musicbrainz artist id
	GroupByAlbumName  - Group by Album Name
	SortByReleaseDate - Sort by release date (asc|desc)
	Page              - Define the page number for paginated results
	PageSize          - Define the page size for paginated results. Range is 1 to 100.
*/
func (client *Client) GetArtistAlbums(ctx context.Context, params ...mxmParams.Param) ([]*Album, error) {

	var albumsData albumList

	err := client.get(ctx, "artist.albums.get", &albumsData, params...)

	if err != nil {
		return nil, err
	}

	var albums []*Album

	for _, album := range albumsData.AlbumList {
		albums = append(albums, &album.AlbumData)
	}

	return albums, nil

}

/*
Get a list of artists somehow related to a given one.

Parameters:

	ArtistID          - Musixmatch artist id
	ArtistMbID        - Musicbrainz artist id
	Page              - Define the page number for paginated results
	PageSize          - Define the page size for paginated results. Range is 1 to 100.
*/
func (client *Client) GetRelatedArtists(ctx context.Context, params ...mxmParams.Param) ([]*Artist, error) {
	var artistData artistList

	err := client.get(ctx, "artist.related.get", &artistData, params...)

	if err != nil {
		return nil, err
	}

	var artists []*Artist

	for _, artist := range artistData.ArtistList {
		artists = append(artists, &artist.ArtistData)
	}

	return artists, nil

}
