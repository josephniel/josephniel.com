package main

import (
	"log"

	"github.com/josephniel/portfolio/config"
	"github.com/josephniel/portfolio/router/https"
)

func main() {
	conf := config.Load()
	router := https.NewRouter(conf)

	if err := router.Serve(); err != nil {
		log.Fatalf("Cannot serve, error found: %e", err)
	}
}
