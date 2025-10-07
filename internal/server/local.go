package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/evanwiseman/ppss-server/internal/config"
	"github.com/evanwiseman/ppss-server/internal/database"
	"github.com/evanwiseman/ppss-server/internal/models"
	"github.com/google/uuid"
	"github.com/lib/pq"
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
	created, err := s.Queries.CreateDevice(r.Context(), database.CreateDeviceParams{
		Name:       d.Name,
		IpAddress:  d.IpAddress,
		DeviceType: d.DeviceType,
	})
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" { // unique_violation
				http.Error(
					w,
					fmt.Sprintf("device with id '%v' already exists", d.ID),
					http.StatusConflict,
				)
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

func (s *LocalServer) DeleteDeviceByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get the device ID
	deviceID, err := uuid.Parse(r.PathValue("deviceID"))
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("invalid id: %v", err),
			http.StatusBadRequest,
		)
		return
	}

	// Delete the device with the matching ID
	err = s.Queries.DeleteDeviceByID(r.Context(), deviceID)
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("device not found: %v", err),
			http.StatusNotFound,
		)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
