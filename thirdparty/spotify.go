package thirdparty

import (
	"context"
	"encoding/json"
	"os"

	"github.com/rafaelbmateus/go-bootstrap/entity"
	"golang.org/x/oauth2/clientcredentials"
)

const tokenURL = "https://accounts.spotify.com/api/token"

type Spotify struct {
	TokenURL     string
	ClientID     string
	ClientSecret string
}

func NewSpotify() *Spotify {
	return &Spotify{
		TokenURL:     tokenURL,
		ClientID:     os.Getenv("SPOTIFY_CLIENT_ID"),
		ClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
	}
}

func (s *Spotify) GetPlaylistByCategory(ctx context.Context, url string) (*entity.PlaylistsRoot, error) {
	// var client = &http.Client{Timeout: 10 * time.Second}
	clientConfig := clientcredentials.Config{
		ClientID:     s.ClientID,
		ClientSecret: s.ClientSecret,
		TokenURL:     s.TokenURL,
	}

	client := clientConfig.Client(ctx)
	r, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	var playlist *entity.PlaylistsRoot
	err = json.NewDecoder(r.Body).Decode(&playlist)
	if err != nil {
		return nil, err
	}

	return playlist, nil
}

func (s *Spotify) GetPlaylistTracks(ctx context.Context, url string) (*entity.Tracks, error) {
	// var client = &http.Client{Timeout: 10 * time.Second}
	clientConfig := clientcredentials.Config{
		ClientID:     s.ClientID,
		ClientSecret: s.ClientSecret,
		TokenURL:     s.TokenURL,
	}

	client := clientConfig.Client(ctx)
	r, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	var tracks *entity.Tracks
	err = json.NewDecoder(r.Body).Decode(&tracks)
	if err != nil {
		return nil, err
	}

	return tracks, nil
}
