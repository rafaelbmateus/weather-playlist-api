package thirdparty_test

import (
	"context"
	"testing"

	"github.com/rafaelbmateus/go-bootstrap/thirdparty"
	"github.com/stretchr/testify/assert"
)

func TestGetPlaylistByCategory(t *testing.T) {
	type test struct {
		url string
	}

	tests := []test{
		{
			url: "https://api.spotify.com/v1/browse/categories/party/playlists?offset=0&limit=1",
		},
		{
			url: "https://api.spotify.com/v1/browse/categories/pop/playlists?offset=0&limit=1",
		},
		{
			url: "https://api.spotify.com/v1/browse/categories/rock/playlists?offset=0&limit=1",
		},
		{
			url: "https://api.spotify.com/v1/browse/categories/classical/playlists?offset=0&limit=1",
		},
	}

	for _, tc := range tests {
		ow := thirdparty.NewSpotify()
		playlist, _ := ow.GetPlaylistByCategory(context.Background(), tc.url)
		assert.Equal(t, playlist.Playlists.URL, tc.url)
		assert.NotNil(t, playlist)
	}
}

func TestGetPlaylistTracks(t *testing.T) {
	type test struct {
		url string
	}

	tests := []test{
		{
			url: "https://api.spotify.com/v1/playlists/37i9dQZF1DX8mBRYewE6or/tracks",
		},
		{
			url: "https://api.spotify.com/v1/playlists/37i9dQZF1DWUa8ZRTfalHk/tracks",
		},
		{
			url: "https://api.spotify.com/v1/playlists/37i9dQZF1DXa8n42306eJB/tracks",
		},
	}

	for _, tc := range tests {
		ow := thirdparty.NewSpotify()
		tracks, _ := ow.GetPlaylistTracks(context.Background(), tc.url)
		assert.NotNil(t, tracks)
	}
}
