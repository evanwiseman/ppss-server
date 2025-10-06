package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/evanwiseman/ppss-server/internal/config"
	"github.com/evanwiseman/ppss-server/internal/database/local"
	"github.com/evanwiseman/ppss-server/internal/models"
	"github.com/lib/pq"
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

	// Parse JSON request
	var d models.Device
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&d)
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("invalid JSON: %v", err),
			http.StatusBadRequest,
		)
		return
	}

	// Create entry in database
	created, err := s.Queries.CreateDevice(r.Context(), local.CreateDeviceParams{
		SerialNumber: d.SerialNumber,
		Name:         d.Name,
		IpAddress:    d.IpAddress,
		DeviceType:   d.DeviceType,
	})
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" { // unique_violation
				http.Error(w, fmt.Sprintf("device with serial_number '%s' already exists", d.SerialNumber), http.StatusConflict)
				return
			}
		}
		http.Error(w, fmt.Sprintf("unable to create device: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with created
	resp := models.DB2Device(created)
	data, err := json.Marshal(resp)
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("unable to pack response: %v", err),
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}
