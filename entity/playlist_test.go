package entity_test

import (
	"testing"

	"github.com/rafaelbmateus/go-bootstrap/entity"
	"github.com/stretchr/testify/assert"
)

func TestPlaylist(t *testing.T) {
	r := entity.PlaylistsRoot{
		Playlists: entity.Playlists{
			Items: []entity.Playlist{
				{
					ID:            "123",
					Collaborative: true,
					Name:          "Playlist name",
					Description:   "Playlist description",
					URL:           "https://api.spotify.com/v1/playlists/37i9dQZF1DWUa8ZRTfalHk",
				},
			}},
	}

	assert.Equal(t, len(r.Playlists.Items), 1)
	assert.Equal(t, r.Playlists.Items[0].ID, "123")
	assert.Equal(t, r.Playlists.Items[0].Collaborative, true)
	assert.Equal(t, r.Playlists.Items[0].Name, "Playlist name")
	assert.Equal(t, r.Playlists.Items[0].Description, "Playlist description")
	assert.Equal(t, r.Playlists.Items[0].URL, "https://api.spotify.com/v1/playlists/37i9dQZF1DWUa8ZRTfalHk")
}
