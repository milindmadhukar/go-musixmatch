package gomusixmatch

import (
	"context"
)

// Get the list of the music genres of Musixmatch's catalogue.
func (client *Client) GetMusicGenres(ctx context.Context) (*[]MusicGenre, error) {
	var genresData musicGenreList

	err := client.get(ctx, "music.genres.get", &genresData)

	if err != nil {
		return nil, err
	}

	var genres []MusicGenre

	for _, genre := range genresData.MusicGenreList {
		genres = append(genres, genre.MusicGenreData)
	}

	return &genres, nil

}
