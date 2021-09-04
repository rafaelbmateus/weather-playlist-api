package entity_test

import (
	"testing"

	"github.com/rafaelbmateus/go-bootstrap/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewCity(t *testing.T) {
	l, err := entity.NewLocation("Campinas", "-47.0608", "-22.9056")
	assert.Nil(t, err)
	assert.Equal(t, l.City, "Campinas")
	assert.Equal(t, l.Lat, "-47.0608")
	assert.Equal(t, l.Lon, "-22.9056")
}

func TestCityValidate(t *testing.T) {
	type test struct {
		city string
		lon  string
		lat  string
		want error
	}

	tests := []test{
		{
			city: "Campinas",
			want: nil,
		},
		{
			city: "Brasília",
			want: nil,
		},
		{
			city: "Manaus",
			want: nil,
		},
		{
			city: "Fortaleza",
			lon:  "-38.5247",
			lat:  "-3.7227",
			want: nil,
		},
		{
			city: "Salvador",
			lon:  "-88.9167",
			lat:  "13.8333",
			want: nil,
		},
		{
			city: "",
			lon:  "-47.9297",
			lat:  "-15.7797",
			want: nil,
		},
		{
			lon:  "1",
			lat:  "",
			want: entity.ErrInvalidEntity,
		},
		{
			lon:  "",
			lat:  "1",
			want: entity.ErrInvalidEntity,
		},
		{
			lon:  "",
			lat:  "",
			want: entity.ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		_, err := entity.NewLocation(tc.city, tc.lat, tc.lon)
		assert.Equal(t, err, tc.want)
	}
}
