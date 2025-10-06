package config

import (
	"os"
)

type Config struct {
	LocalAddr  string
	PublicAddr string
	DBURL      string
	Platform   string
}

func Load() (*Config, error) {
	cfg := &Config{
		LocalAddr:  os.Getenv("LOCAL_ADDR"),
		PublicAddr: os.Getenv("PUBLIC_ADDR"),
		DBURL:      os.Getenv("DB_URL"),
		Platform:   os.Getenv("PLATFORM"),
	}

	return cfg, nil
}
