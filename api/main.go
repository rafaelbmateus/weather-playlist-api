package main

import (
	"fmt"

	"github.com/rafaelbmateus/go-bootstrap/api/handler"
	"github.com/rs/zerolog/log"
)

var (
	Version = "no version provided"
	Commit  = "no commit hash provided"
)

func main() {
	log.Info().Msg(fmt.Sprintf("Starting using version %s and commit %s", Version, Commit))
	handler.Routes()
}
