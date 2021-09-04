package entity_test

import (
	"testing"

	"github.com/rafaelbmateus/go-bootstrap/entity"
	"github.com/stretchr/testify/assert"
)

func TestWeather(t *testing.T) {
	weather := entity.Weather{
		ID:   123,
		Name: "City name",
		Main: struct {
			Temp      float32 `json:"temp"`
			FeelsLike float32 `json:"feels_like"`
			TempMin   float32 `json:"temp_min"`
			TempMax   float32 `json:"temp_max"`
			Pressure  int     `json:"pressure"`
			Humidity  int     `json:"humidity"`
		}{Temp: 10.7, FeelsLike: 11, TempMin: 0, TempMax: 30, Pressure: 20, Humidity: 10}}

	assert.Equal(t, weather.ID, 123)
	assert.Equal(t, weather.Name, "City name")
	assert.Equal(t, weather.Main.Temp, float32(10.7))
	assert.Equal(t, weather.Main.FeelsLike, float32(11))
	assert.Equal(t, weather.Main.TempMin, float32(0))
	assert.Equal(t, weather.Main.TempMax, float32(30))
	assert.Equal(t, weather.Main.Pressure, 20)
	assert.Equal(t, weather.Main.Humidity, 10)
}
