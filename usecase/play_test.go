package usecase_test

import (
	"context"
	"testing"

	"github.com/rafaelbmateus/go-bootstrap/entity"
	"github.com/rafaelbmateus/go-bootstrap/usecase"
	"github.com/stretchr/testify/assert"
)

func TestSearchPlaylist(t *testing.T) {
	type test struct {
		location entity.Location
		want     error
	}

	tests := []test{
		{
			location: entity.Location{City: "Campinas"},
			want:     nil,
		},
		{
			location: entity.Location{City: "Brasília"},
			want:     nil,
		},
		{
			location: entity.Location{City: "Fortaleza"},
			want:     nil,
		},
		{
			location: entity.Location{Lat: "13.8333", Lon: "-88.9167"},
			want:     nil,
		},
		{
			location: entity.Location{Lat: "-3.1019", Lon: "-60.025"},
			want:     nil,
		},
	}

	for _, tc := range tests {
		weather, playlist, tracks, err := usecase.SearchPlaylist(context.Background(), tc.location)
		assert.NotNil(t, weather)
		assert.NotNil(t, playlist)
		assert.NotNil(t, tracks)
		assert.Equal(t, tc.want, err)
	}
}

func TestWeatherQuery(t *testing.T) {
	type test struct {
		location entity.Location
		want     string
	}

	tests := []test{
		{
			location: entity.Location{City: "Campinas"},
			want:     "q=Campinas",
		},
		{
			location: entity.Location{City: "Brasília"},
			want:     "q=Brasília",
		},
		{
			location: entity.Location{Lat: "13.8333", Lon: "-88.9167"},
			want:     "lat=13.8333&lon=-88.9167",
		},
		{
			location: entity.Location{Lat: "-3.1019", Lon: "-60.025"},
			want:     "lat=-3.1019&lon=-60.025",
		},
	}

	for _, tc := range tests {
		query := usecase.WeatherQuery(tc.location)
		assert.Equal(t, tc.want, query)
	}
}

func TestChoiceCategory(t *testing.T) {
	type test struct {
		temperature float32
		want        string
	}

	tests := []test{
		{
			temperature: 0,
			want:        "classical",
		},
		{
			temperature: 9.9999,
			want:        "classical",
		},
		{
			temperature: 10,
			want:        "rock",
		},
		{
			temperature: 14.9,
			want:        "rock",
		},
		{
			temperature: 15,
			want:        "pop",
		},
		{
			temperature: 30.6,
			want:        "pop",
		},
		{
			temperature: 31,
			want:        "party",
		},
		{
			temperature: 99,
			want:        "party",
		},
		{
			temperature: -3827.210,
			want:        "classical",
		},
	}

	for _, tc := range tests {
		query := usecase.ChoiceCategory(tc.temperature)
		assert.Equal(t, tc.want, query)
	}
}
