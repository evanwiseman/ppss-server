package server

import (
	"github.com/evanwiseman/ppss-server/internal/config"
	"github.com/evanwiseman/ppss-server/internal/database"
)

type LocalServer struct {
	Cfg     *config.Config
	Queries *database.Queries
}

func NewLocalServer(cfg *config.Config, queries *database.Queries) (*LocalServer, error) {
	return &LocalServer{
		Cfg:     cfg,
		Queries: queries,
	}, nil
}
