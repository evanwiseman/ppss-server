package config

import (
	"os"
)

type Config struct {
	LocalAddress  string
	PublicAddress string
	DatabaseURL   string
	Platform      string
}

func Load() (*Config, error) {
	cfg := &Config{
		LocalAddress:  os.Getenv("LOCAL_ADDRESS"),
		PublicAddress: os.Getenv("PUBLIC_ADDRESS"),
		DatabaseURL:   os.Getenv("DATABASE_URL"),
		Platform:      os.Getenv("PLATFORM"),
	}

	return cfg, nil
}
