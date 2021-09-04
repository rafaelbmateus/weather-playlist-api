package thirdparty

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/rafaelbmateus/go-bootstrap/entity"
)

type OpenWeather struct {
	URL   string
	AppID string
}

func NewOpenWeather() *OpenWeather {
	return &OpenWeather{
		AppID: "b77e07f479efe92156376a8b07640ced",
	}
}

// API Access current weather data for any location - https://openweathermap.org/current
func (ow *OpenWeather) API(ctx context.Context, url string) (*entity.Weather, error) {
	var client = &http.Client{Timeout: 10 * time.Second}
	r, err := client.Get(url + "&appid=" + ow.AppID)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	var weather *entity.Weather
	err = json.NewDecoder(r.Body).Decode(&weather)

	if weather.ID == 0 {
		return nil, entity.ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	return weather, nil
}
