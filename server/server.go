package main

import (
	"log"

	"github.com/josephniel/josephniel.com/server/config"
	"github.com/josephniel/josephniel.com/server/router/https"
)

func main() {
	conf := config.Load()
	router := https.NewRouter(conf)

	if err := router.Serve(); err != nil {
		log.Fatalf("Cannot serve, error found: %e", err)
	}
}
