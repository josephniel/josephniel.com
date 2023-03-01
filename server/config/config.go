package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Host string
	Port int
}

func Load() Config {
	// will not overwrite existing envs
	_ = godotenv.Load(".env")

	port, err := strconv.ParseInt(os.Getenv("PORT"), 0, 0)
	if err != nil {
		panic("PORT misconfigured")
	}

	return Config{
		Host: os.Getenv("HOST"),
		Port: int(port),
	}
}
