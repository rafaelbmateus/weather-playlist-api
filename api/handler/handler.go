package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Routes() {
	router := gin.Default()
	router.GET("/", Home)
	router.GET("/playlist", Play)
	log.Info().Msg("service is up")
	err := router.Run(":3000")
	if err != nil {
		log.Fatal().Msg("error on run server")
	}
}

func Home(c *gin.Context) {
	c.JSON(200, gin.H{
		"title": "Welcome to the Weather Playlist API",
		"samples": []string{
			fmt.Sprintf("http://%s/playlist?city=campinas", c.Request.Host),
			fmt.Sprintf("http://%s/playlist?city=salvador", c.Request.Host),
			fmt.Sprintf("http://%s/playlist?city=fortaleza", c.Request.Host),
			fmt.Sprintf("http://%s/playlist?city=florianópolis", c.Request.Host),
			fmt.Sprintf("http://%s/playlist?lat=-3.1019&lon=-60.025", c.Request.Host),
			fmt.Sprintf("http://%s/playlist?lat=13.8333&lon=-88.9167", c.Request.Host),
		},
	})
}
