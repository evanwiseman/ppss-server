package server

import (
	"database/sql"

	"github.com/evanwiseman/ppss-server/internal/config"
)

type BaseServer struct {
	CFG *config.Config
	DB  *sql.DB
}
