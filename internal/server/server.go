package server

import (
	"database/sql"

	"github.com/evanwiseman/ppss-server/internal/config"
)

type Server struct {
	CFG *config.Config
	DB  *sql.DB
}

func NewServer(cfg *config.Config, dbURL string) (*Server, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	return &Server{
		CFG: cfg,
		DB:  db,
	}, nil
}
