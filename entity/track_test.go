package entity_test

import (
	"testing"

	"github.com/rafaelbmateus/go-bootstrap/entity"
	"github.com/stretchr/testify/assert"
)

func TestTrack(t *testing.T) {
	tracks := entity.Tracks{
		Items: []entity.TrackItem{
			{
				Track: entity.Track{
					ID:         "123",
					Name:       "Track name",
					URL:        "https://",
					PreviewURL: "https://",
				},
			},
		},
	}

	assert.Equal(t, len(tracks.Items), 1)
	assert.Equal(t, tracks.Items[0].Track.ID, "123")
	assert.Equal(t, tracks.Items[0].Track.Name, "Track name")
	assert.Equal(t, tracks.Items[0].Track.URL, "https://")
	assert.Equal(t, tracks.Items[0].Track.PreviewURL, "https://")
}
