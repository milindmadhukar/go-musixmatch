package gomusixmatch

import (
	"context"
	"fmt"
)

type musicGenresGet struct {
	MusicGenreList []struct {
		MusicGenreData MusicGenre `json:"music_genre"`
	} `json:"music_genre_list"`
	Url string `json:"url,omitempty"`
}

func (client *Client) GetMusicGenres(ctx context.Context) (*musicGenresGet, error) {
	url := fmt.Sprintf("%smusic.genres.get?apikey=%s",
		client.baseURL,
		client.apiKey)

	var get_genres musicGenresGet

	err := client.get(ctx, url, &get_genres)

	if err != nil {
		return nil, err
	}

	get_genres.Url = url

	return &get_genres, nil

}
