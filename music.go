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

// Get the list of the music genres of Musixmatch's catalogue.
func (client *Client) GetMusicGenres(ctx context.Context) (*musicGenresGet, error) {
	url := fmt.Sprintf("%smusic.genres.get?apikey=%s",
		client.baseURL,
		client.apiKey)

	var getGenres musicGenresGet

	err := client.get(ctx, url, &getGenres)

	if err != nil {
		return nil, err
	}

	getGenres.Url = url

	return &getGenres, nil

}
