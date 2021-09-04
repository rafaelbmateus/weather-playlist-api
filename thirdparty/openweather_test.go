package thirdparty_test

import (
	"context"
	"testing"

	"github.com/rafaelbmateus/go-bootstrap/thirdparty"
	"github.com/stretchr/testify/assert"
)

func TestWeatherAPI(t *testing.T) {
	type test struct {
		url string
	}

	tests := []test{
		{
			url: "http://api.openweathermap.org/data/2.5/weather?q=Campinas&units=metric",
		},
		{
			url: "http://api.openweathermap.org/data/2.5/weather?q=Salvador&units=metric",
		},
		{
			url: "http://api.openweathermap.org/data/2.5/weather?q=Brasília&units=metric",
		},
		{
			url: "http://api.openweathermap.org/data/2.5/weather?q=Fortaleza&units=metric",
		},
		{
			url: "http://api.openweathermap.org/data/2.5/weather?q=Manaus&units=metric",
		},
	}

	for _, tc := range tests {
		ow := thirdparty.NewOpenWeather()
		w, _ := ow.API(context.Background(), tc.url)

		assert.NotNil(t, w.Name)
		assert.NotNil(t, w.Main.Temp)
	}
}
