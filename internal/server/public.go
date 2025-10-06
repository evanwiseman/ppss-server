package server

import (
	"database/sql"

	"github.com/evanwiseman/ppss-server/internal/config"
	"github.com/evanwiseman/ppss-server/internal/database/public"
)

type PublicServer struct {
	BaseServer
	Queries *public.Queries
}

func NewPublicServer(cfg *config.Config, dbURL string) (*PublicServer, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	return &PublicServer{
		BaseServer: BaseServer{CFG: cfg},
		Queries:    public.New(db),
	}, nil
}
