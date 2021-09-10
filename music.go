package gomusixmatch

import (
	"context"
	"fmt"
)

type musicGenresGet struct {
	Message struct {
		Header struct {
			StatusCode  int     `json:"status_code"`
			ExecuteTime float64 `json:"execute_time"`
		} `json:"header"`
		Body struct {
			MusicGenreList []struct {
				MusicGenreData MusicGenre `json:"music_genre"`
			} `json:"music_genre_list"`
		} `json:"body"`
	} `json:"message"`
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

	return &get_genres, nil

}
