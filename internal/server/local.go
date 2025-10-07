package server

import (
	"encoding/json"
	"errors"
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
		models.RespondWithError(
			w,
			http.StatusBadRequest,
			"invalid JSON",
			err,
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
				models.RespondWithError(
					w,
					http.StatusConflict,
					"device already exists",
					err,
				)
				return
			}
		}
		models.RespondWithError(
			w,
			http.StatusInternalServerError,
			"unable to create device",
			err,
		)
		return
	}

	// Respond with created
	resp := models.DB2Device(created)
	models.RespondWithJSON(w, http.StatusCreated, resp)
}

func (s *LocalServer) DeleteDeviceByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get the device ID
	deviceID, err := uuid.Parse(r.PathValue("deviceID"))
	if err != nil {
		models.RespondWithError(
			w,
			http.StatusBadRequest,
			"invalid id",
			err,
		)
		return
	}

	// Check the device exists
	_, err = s.Queries.GetDeviceByID(r.Context(), deviceID)
	if err != nil {
		models.RespondWithError(
			w,
			http.StatusNotFound,
			"device not found",
			err,
		)
		return
	}

	// Delete the device with the matching ID
	err = s.Queries.DeleteDeviceByID(r.Context(), deviceID)
	if err != nil {
		models.RespondWithError(
			w,
			http.StatusInternalServerError,
			"unable to delete device",
			err,
		)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *LocalServer) GetDevicesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get all devices
	rows, err := s.Queries.GetDevices(r.Context())
	if err != nil {
		models.RespondWithError(
			w,
			http.StatusNotFound,
			"no devices found",
			err,
		)
	}

	// Send the response
	var resp []models.Device
	for _, row := range rows {
		resp = append(resp, models.DB2Device(row))
	}
	models.RespondWithJSON(w, http.StatusOK, resp)
}

func (s *LocalServer) GetDeviceByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get the device ID
	deviceID, err := uuid.Parse(r.PathValue("deviceID"))
	if err != nil {
		models.RespondWithError(
			w,
			http.StatusBadRequest,
			"invalid id",
			err,
		)
		return
	}

	// Get the device
	row, err := s.Queries.GetDeviceByID(r.Context(), deviceID)
	if err != nil {
		models.RespondWithError(
			w,
			http.StatusNotFound,
			"device not found",
			err,
		)
		return
	}

	// Send the response
	resp := models.DB2Device(row)
	models.RespondWithJSON(w, http.StatusOK, resp)
}

func (s *LocalServer) ResetDevicesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Check that platform is in dev
	// TODO: Check if user is authenticated to reset
	if s.Cfg.Platform != "dev" {
		models.RespondWithError(
			w,
			http.StatusUnauthorized,
			"not authorized to reset devices",
			errors.New("database not on dev platform"),
		)
		return
	}

	// Reset devices in the database
	err := s.Queries.ResetDevices(r.Context())
	if err != nil {
		models.RespondWithError(
			w,
			http.StatusInternalServerError,
			"unable to reset devices",
			err,
		)
	}

	w.WriteHeader(http.StatusNoContent)
}
