package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafaelbmateus/go-bootstrap/entity"
	"github.com/rafaelbmateus/go-bootstrap/usecase"
	"github.com/rs/zerolog/log"
)

func Play(c *gin.Context) {
	city := c.Query("city")
	lat := c.Query("lat")
	lon := c.Query("lon")

	log.Info().Msgf("request received with params city=%s lat=%s lon=%s", city, lat, lon)

	location, err := entity.NewLocation(city, lat, lon)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "you need to send city or lat and lon params"})
		return
	}

	weather, playlist, tracks, err := usecase.SearchPlaylist(context.Background(), *location)
	if err != nil {
		if err == entity.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		log.Info().Msgf("error on search playlist - %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, struct {
		Weather  entity.Weather  `json:"weather"`
		Playlist entity.Playlist `json:"playlist"`
		Tracks   entity.Tracks   `json:"tracks"`
	}{
		Weather:  *weather,
		Playlist: *playlist,
		Tracks:   *tracks,
	})
}
