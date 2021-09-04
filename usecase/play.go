package usecase

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/rafaelbmateus/go-bootstrap/entity"
	"github.com/rafaelbmateus/go-bootstrap/thirdparty"

	"github.com/rs/zerolog/log"
)

var (
	spotifyCategoryURL       = "https://api.spotify.com/v1/browse/categories/%s/playlists?offset=0&limit=10"
	spotifyPlaylistTracksURL = "https://api.spotify.com/v1/playlists/%s/tracks"
	openWeatherURL           = "http://api.openweathermap.org/data/2.5/weather?%s&units=metric"
)

// SearchPlaylist search a playlist of the location depending on the weather.
func SearchPlaylist(ctx context.Context, location entity.Location) (*entity.Weather, *entity.Playlist, *entity.Tracks, error) {
	log.Info().Msg(fmt.Sprintf("searching a playlist for location %+v", location))
	ow := thirdparty.NewOpenWeather()
	weather, err := ow.API(ctx, fmt.Sprintf(openWeatherURL, WeatherQuery(location)))
	log.Info().Msg("query open weather: " + fmt.Sprintf(openWeatherURL, WeatherQuery(location)))
	if err != nil {
		return nil, nil, nil, err
	}

	category := ChoiceCategory(weather.Main.Temp)
	log.Info().Msg(fmt.Sprintf("%+v - %.0f°C, %s is a good song to play",
		location, weather.Main.Temp, category))

	spotify := thirdparty.NewSpotify()
	playlist, err := spotify.GetPlaylistByCategory(ctx, fmt.Sprintf(spotifyCategoryURL, category))
	if err != nil {
		return nil, nil, nil, err
	}

	choicePlaylist := playlist.Playlists.Items[rand.Intn(len(playlist.Playlists.Items))]
	log.Info().Msg(fmt.Sprintf("I found %d %s playlists. I chose the %+v",
		len(playlist.Playlists.Items), category, choicePlaylist))

	tracks, err := spotify.GetPlaylistTracks(ctx,
		fmt.Sprintf(spotifyPlaylistTracksURL, choicePlaylist.ID))
	if err != nil {
		return nil, nil, nil, err
	}

	return weather, &choicePlaylist, tracks, nil
}

// WeatherQuery make open weather query by city name or location.
func WeatherQuery(location entity.Location) string {
	if location.City != "" {
		return "q=" + location.City
	}

	return "lat=" + location.Lat + "&lon=" + location.Lon
}

// ChoiceCategory choice a music category by temperature.
func ChoiceCategory(temperature float32) string {
	if temperature < 10 {
		return "classical"
	}

	if temperature >= 10 && temperature < 15 {
		return "rock"
	}

	if temperature >= 15 && temperature < 31 {
		return "pop"
	}

	return "party"
}
