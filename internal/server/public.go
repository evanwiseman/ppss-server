package server

import (
	"github.com/evanwiseman/ppss-server/internal/config"
	"github.com/evanwiseman/ppss-server/internal/database"
)

type PublicServer struct {
	Cfg     *config.Config
	Queries *database.Queries
}

func NewPublicServer(cfg *config.Config, queries *database.Queries) (*PublicServer, error) {
	return &PublicServer{
		Cfg:     cfg,
		Queries: queries,
	}, nil
}
