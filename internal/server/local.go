package server

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/evanwiseman/ppss-server/internal/config"
	"github.com/evanwiseman/ppss-server/internal/database/local"
	"github.com/evanwiseman/ppss-server/internal/models"
)

type LocalServer struct {
	BaseServer
	Queries *local.Queries
}

func NewLocalServer(cfg *config.Config, dbURL string) (*LocalServer, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	return &LocalServer{
		BaseServer: BaseServer{CFG: cfg},
		Queries:    local.New(db),
	}, nil
}

func (s *LocalServer) PostDeviceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var d models.Device
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&d)
	if err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

}
